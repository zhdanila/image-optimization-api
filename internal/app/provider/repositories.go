package provider

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/samber/do/v2"
	"image-optimization-api/internal/app"
	"image-optimization-api/internal/repository"
)

func ProvideImageRepository(inj do.Injector) (*repository.Image, error) {
	cnf := do.MustInvoke[*app.Config](inj)

	return repository.NewImage(do.MustInvoke[*s3.S3](inj), cnf.S3Bucket), nil
}
