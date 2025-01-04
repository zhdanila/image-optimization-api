package repository

import (
	"context"
	"github.com/aws/aws-sdk-go/service/s3"
	"image-optimization-api/internal/domain/image"
	"image-optimization-api/pkg/bind"
	"image-optimization-api/pkg/schema"
)

type Image struct {
	schema.Repository[image.Image]
}

func NewImage(db *s3.S3) *Image {
	return &Image{
		Repository: schema.NewRepository(db, image.Image{}),
	}
}

func (r *Image) UploadImages(ctx context.Context, images []bind.UploadedFile) error {
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
