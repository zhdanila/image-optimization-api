package repository

import (
	"image-optimization-api/internal/domain/image"
	"image-optimization-api/pkg/schema"
)

type Image struct {
	schema.Repository[image.Image]
}

func NewImage() *Image {
	return &Image{
		Repository: schema.NewRepository(nil, image.Image{}),
	}
}
