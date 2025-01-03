package provider

import (
	"github.com/samber/do/v2"
	"image-optimization-api/internal/repository"
)

func ProvideImageRepository(inj do.Injector) (*repository.Image, error) {
	return repository.NewImage(), nil
}
