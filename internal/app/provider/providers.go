package provider

import (
	"github.com/samber/do/v2"
	"image-optimization-api/internal/app"
	"image-optimization-api/internal/app/interface/http/website"
	"image-optimization-api/pkg/server"
)

func ProvideConfig(_ do.Injector) (*app.Config, error) {
	return app.NewConfig()
}

func ProvideWebsiteServer(inj do.Injector) (*server.Server, error) {
	cnf := do.MustInvoke[*app.Config](inj)

	return website.NewServer(cnf, inj), nil
}
