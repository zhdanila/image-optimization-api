package repository

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"image-optimization-api/internal/domain/image"
	"image-optimization-api/pkg/bind"
)

type Image struct {
	db         *s3.S3
	bucketName string
	region     string
}

func NewImage(db *s3.S3, bucketName, region string) *Image {
	return &Image{
		db:         db,
		bucketName: bucketName,
		region:     region,
	}
}

func (r *Image) UploadImages(ctx context.Context, images []bind.UploadedFile) error {
	for _, img := range images {
		_, err := r.db.PutObjectWithContext(ctx, &s3.PutObjectInput{
			Bucket:      aws.String(r.bucketName),
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

func (r *Image) GetImage(ctx context.Context, id string) (*image.ImageInfo, error) {
	headObjectInput := &s3.HeadObjectInput{
		Bucket: aws.String(r.bucketName),
		Key:    aws.String(id),
	}

	_, err := r.db.HeadObjectWithContext(ctx, headObjectInput)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", r.bucketName, r.region, id)

	return &image.ImageInfo{
		Key: id,
		URL: url,
	}, nil
}

func (r *Image) ListImages(ctx context.Context) ([]image.ImageInfo, error) {
	var images []image.ImageInfo

	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(r.bucketName),
	}

	result, err := r.db.ListObjectsV2WithContext(ctx, input)
	if err != nil {
		return nil, err
	}

	for _, item := range result.Contents {
		url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", r.bucketName, r.region, *item.Key)
		images = append(images, image.ImageInfo{
			Key: *item.Key,
			URL: url,
		})
	}

	return images, err
}
