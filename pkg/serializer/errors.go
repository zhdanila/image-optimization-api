package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

// Error is an error that can be returned by a gRPC service. It contains gRPC status code, a message, and an
// internal error if there is one.
type Error struct {
	Code     int         `json:"code"`
	Message  interface{} `json:"message"`
	Internal error       `json:"-"`
}

// Error implements the error interface for the Error struct.
func (e *Error) Error() string {
	return fmt.Sprintf("code: %d, message: %v", e.Code, e.Message)
}

// NewError returns a new Error with the given internal error.
func NewError(code int, message ...any) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// SetInternal sets the internal error of the Error.
func (e *Error) SetInternal(err error) *Error {
	e.Internal = err
	return e
}

// HTTPStatus returns the HTTP status text corresponding to the code.
func (e *Error) HTTPStatus() string {
	return http.StatusText(e.Code)
}

// HTTPErrorHandler handles errors in the HTTP response.
func HTTPErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	var he *Error

	switch e := err.(type) {
	case *Error:
		he = e
		if he.Internal != nil {
			var herr *Error
			if errors.As(he.Internal, &herr) {
				he = herr
			}
		}
	case *echo.HTTPError:
		he = NewError(e.Code, e.Message)
		if he.Internal != nil {
			var herr *Error
			if errors.As(he.Internal, &herr) {
				he = herr
			}
		}
	default:
		he = NewError(
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
		)
	}

	code := he.Code
	message := he.Message

	switch m := he.Message.(type) {
	case string:
		zap.L().Error(fmt.Sprintf("HTTP error: %v", err))
		message = echo.Map{"message": m}
	case json.Marshaler:
		// do nothing - this type knows how to format itself to JSON
	case validator.ValidationErrors:
		errors := make([]string, len(m))
		for i, fieldError := range m {
			errors[i] = fmt.Sprintf("field '%s' failed validation with tag '%s' and value '%v'",
				fieldError.Field(), fieldError.Tag(), fieldError.Value())
		}
		message = echo.Map{
			"message": `Validation failed`,
			"errors":  errors,
		}
	case error:
		zap.L().Error(fmt.Sprintf("HTTP error: %v", err))
		message = echo.Map{"message": m.Error()}
	}

	// Send response
	if c.Request().Method == http.MethodHead {
		err = c.NoContent(he.Code)
	} else {
		err = c.JSON(code, message)
	}
	if err != nil {
		zap.L().Error("Error sending response", zap.Error(err))
	}
}
