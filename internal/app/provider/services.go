package provider

import (
	"github.com/samber/do/v2"
	"image-optimization-api/internal/repository"
	"image-optimization-api/internal/service/image"
)

func ProvideImageService(inj do.Injector) (*image.Service, error) {
	return image.NewService(
		do.MustInvoke[*repository.Image](inj),
	), nil
}
