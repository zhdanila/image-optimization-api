package image

import (
	"image-optimization-api/internal/repository"
)

func NewService(
	imageRepo *repository.Image,
) *Service {
	return &Service{
		imageRepo: imageRepo,
	}
}

type Service struct {
	imageRepo *repository.Image
}
