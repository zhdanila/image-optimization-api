package bind

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

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

func FromMultipartForm(arg ...any) func(c echo.Context, obj any) error {
	return func(c echo.Context, obj any) error {
		if arg == nil {
			return nil
		}
		if len(arg)%2 != 0 {
			return fmt.Errorf("invalid number of arguments")
		}

		for i := 0; i < len(arg); i += 2 {
			key := arg[i].(string)
			form, err := c.MultipartForm()
			if err != nil {
				return err
			}
			files := form.File[key]

			switch dst := arg[i+1].(type) {
			case []UploadedFile:
				maxSize := len(dst)
				for i, file := range files {
					src, err := file.Open()
					if err != nil {
						return err
					}
					defer src.Close()

					fileBytes, err := io.ReadAll(src)
					if err != nil {
						return err
					}

					dst[i] = UploadedFile{
						FileName:    file.Filename,
						ContentType: file.Header.Get("Content-Type"),
						Size:        file.Size,
						Src:         fileBytes,
						Tag:         key,
					}
					if i == maxSize-1 {
						break
					}
				}
			case *UploadedFile:
				for _, file := range files {
					src, err := file.Open()
					if err != nil {
						return err
					}
					defer src.Close()

					fileBytes, err := io.ReadAll(src)
					if err != nil {
						return err
					}

					if dst == nil {
						return fmt.Errorf("invalid destination")
					}
					dst.FileName = file.Filename
					dst.ContentType = file.Header.Get("Content-Type")
					dst.Size = file.Size
					dst.Src = fileBytes
				}
			}
		}

		return nil
	}
}

type UploadedFile struct {
	FileName    string `json:"filename"`
	ContentType string `form:"content_type"`
	Size        int64  `form:"size"`
	Src         []byte `form:"src"`
	Tag         string `form:"tag"`
}
