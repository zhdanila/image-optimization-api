package repository

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"image-optimization-api/pkg/bind"
)

type Image struct {
	db         *s3.S3
	bucketName string
}

func NewImage(db *s3.S3, bucketName string) *Image {
	return &Image{
		db:         db,
		bucketName: bucketName,
	}
}

func (r *Image) UploadImages(ctx context.Context, images []bind.UploadedFile) error {
	for _, img := range images {
		_, err := r.db.PutObjectWithContext(ctx, &s3.PutObjectInput{
			Bucket:      aws.String("amazon-images-storage"),
			Key:         aws.String(img.FileName),
			Body:        bytes.NewReader(img.Src),
			ContentType: aws.String(img.ContentType),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Image) GetImage(ctx context.Context) error {
	var err error

	return err
}

func (r *Image) ListImages(ctx context.Context) error {
	var err error

	return err
}
