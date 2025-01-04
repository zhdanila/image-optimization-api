package repository

import (
	"context"
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

func (r *Image) SaveImage(ctx context.Context) error {
	var err error

	return err
}

func (r *Image) GetImage(ctx context.Context) error {
	var err error

	return err
}

func (r *Image) ListImages(ctx context.Context) error {
	var err error

	return err
}
