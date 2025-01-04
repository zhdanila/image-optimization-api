package website

import (
	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
	"image-optimization-api/internal/app"
	http2 "image-optimization-api/internal/app/interface/http"
	"image-optimization-api/internal/app/interface/http/middleware"
	"image-optimization-api/internal/app/interface/http/website/handler"
	"image-optimization-api/internal/service/image"
	"image-optimization-api/pkg/server"
	"net/http"
)

func NewServer(cnf *app.Config, inj do.Injector) *server.Server {
	server := server.NewServer(cnf, inj)
	server.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"service": "image-optimization-api",
			"message": "Root page do nothing. Please go somewhere else",
		})
	})

	server.Use(middleware.SetCORS())
	server.Use(middleware.Recover())

	server.Validator = http2.CustomValidator()

	nameGroup := server.Group("/api")
	handler.NewImage(
		do.MustInvoke[*image.Service](inj),
	).Register(nameGroup)

	server.GET("/healthz", func(c echo.Context) error {
		return c.JSON(http.StatusOK, `{"status":"OK"}`)
	})

	return server
}
