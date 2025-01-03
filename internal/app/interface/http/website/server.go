package website

import (
	"github.com/labstack/echo/v4"
	"github.com/samber/do/v2"
	"image-optimization-api/internal/app"
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

	nameGroup := server.Group("/")

	handler.NewAuth(
		do.MustInvoke[*image.Service](inj),
	).Register(nameGroup)

	server.GET("/healthz", func(c echo.Context) error {
		return c.JSON(http.StatusOK, `{"status":"OK"}`)
	})

	return server
}
