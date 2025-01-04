package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func Recover() echo.MiddlewareFunc {
	config := middleware.DefaultRecoverConfig
	config.LogErrorFunc = func(c echo.Context, err error, stack []byte) error {
		log.Errorf("PANIC RECOVER: %v %s", err, stack)
		return err
	}

	return middleware.RecoverWithConfig(config)
}

func SetCORS() func(next echo.HandlerFunc) echo.HandlerFunc {
	var origins []string

	origins = []string{
		"http://localhost:8000",
		"http://localhost:63342",
	}

	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     origins,
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
		ExposeHeaders:    []string{echo.HeaderContentType},
	})
}
