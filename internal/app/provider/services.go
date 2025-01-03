package provider

import (
	"github.com/samber/do/v2"
	"github.com/streadway/amqp"
	"image-optimization-api/internal/repository"
	"image-optimization-api/internal/service/compression"
	"image-optimization-api/internal/service/image"
)

func ProvideImageService(inj do.Injector) (*image.Service, error) {
	return image.NewService(
		do.MustInvoke[*repository.Image](inj),
	), nil
}

func ProvideCompressionService(inj do.Injector) (*compression.Service, error) {
	return compression.NewService(
		do.MustInvoke[*amqp.Connection](inj),
		do.MustInvoke[*repository.Image](inj),
	), nil
}
