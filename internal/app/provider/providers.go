package provider

import (
	"github.com/samber/do/v2"
	"image-optimization-api/internal/app"
)

func ProvideConfig(_ do.Injector) (*app.Config, error) {
	return app.NewConfig()
}
