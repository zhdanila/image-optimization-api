package bind

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"mime/multipart"
	"net/http"
	"reflect"
)

const MaxFileSize = 10 << 20

type BindOption func(c echo.Context, obj any) error

func BindValidate(c echo.Context, obj any, opts ...BindOption) error {
	var err error

	err = c.Bind(obj)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error").SetInternal(err)
	}

	if len(opts) > 0 {
		for _, opt := range opts {
			if err = opt(c, obj); err != nil {
				return err
			}
		}
	}

	err = c.Validate(obj)
	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			return echo.NewHTTPError(http.StatusBadRequest, validationErrors)
		}
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	return nil
}

func FromHeaders() func(c echo.Context, obj any) error {
	return func(c echo.Context, obj any) error {
		return (&echo.DefaultBinder{}).BindHeaders(c, obj)
	}
}

func FromQuery() func(c echo.Context, obj any) error {
	return func(c echo.Context, obj any) error {
		return (&echo.DefaultBinder{}).BindQueryParams(c, &obj)
	}
}

type FieldName string

func (f FieldName) String() string {
	return string(f)
}

const (
	FieldNameImage FieldName = "image"
)

func FromMultipartFile(fieldName FieldName) func(c echo.Context, obj any) error {
	return func(c echo.Context, obj any) error {
		err := c.Request().ParseMultipartForm(MaxFileSize)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Failed to process form")
		}

		fileHeader, err := c.FormFile(fieldName.String())
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Failed to retrieve file")
		}

		v := reflect.ValueOf(obj).Elem()
		if v.Kind() != reflect.Struct {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid object type")
		}

		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fieldType := v.Type().Field(i)

			if field.Type() == reflect.TypeOf((*multipart.FileHeader)(nil)) && fieldType.Tag.Get("form") == fieldName.String() {
				field.Set(reflect.ValueOf(fileHeader))
				return nil
			}
		}

		return echo.NewHTTPError(http.StatusBadRequest, "No matching field found for file")
	}
}
