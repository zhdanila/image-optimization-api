package rest

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

var Errs = struct {
	OwnershipCheckFailed *Error
	InvalidArgument      *Error
}{
	OwnershipCheckFailed: NewError(http.StatusPreconditionFailed, "Ownership check failed"),
	InvalidArgument:      NewError(http.StatusBadRequest, "Invalid argument"),
}

// Error is an error that can be returned by a gRPC service. It contains gRPC status code, a message, and an
// internal error if there is one.
//
// Should be used in gRPC services and parsed by the gRPC interceptor.
type Error struct {
	*echo.HTTPError
}

// NewError returns a new Error with the given internal error.
func NewError(code int, message ...any) *Error {
	return &Error{echo.NewHTTPError(code, message...)}
}

// SetInternal sets the internal error of the HTTPError.
func (s *Error) SetInternal(err error) *Error {
	s.Internal = err
	return s
}

func (s *Error) HTTPStatus() string {
	return http.StatusText(s.Code)
}

func HTTPErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	var he *Error

	switch e := err.(type) {
	case *Error:
		he = e
		if he.Internal != nil {
			if herr, ok := he.Internal.(*Error); ok {
				he = herr
			}
		}
	case *echo.HTTPError:
		he = NewError(e.Code, e.Message)
		if he.Internal != nil {
			if herr, ok := he.Internal.(*Error); ok {
				he = herr
			}
		}
	default:
		he = NewError(
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
		)
	}

	// Issue #1426
	code := he.Code
	message := he.Message

	switch m := he.Message.(type) {
	case string:
		zap.L().Error(fmt.Sprintf("HTTP error: %v", err))
		//if e.Debug {
		//	message = echo.Map{"message": m, "error": err.Error()}
		//} else {
		message = echo.Map{"message": m}
		//}
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
	if c.Request().Method == http.MethodHead { // Issue #608
		err = c.NoContent(he.Code)
	} else {
		err = c.JSON(code, message)
	}
	if err != nil {
		//e.Logger.Error(err)
	}
}
