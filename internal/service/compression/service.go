package compression

import (
	"context"
	"fmt"
	"image-optimization-api/internal/repository"
	"image-optimization-api/internal/service/image"
	"image-optimization-api/pkg/imageproc"
)

func NewService(imageRepo *repository.Image) *Service {
	return &Service{
		imageRepo: imageRepo,
	}
}

type Service struct {
	imageRepo *repository.Image
}

func (s *Service) compressImage(ctx context.Context, msg []byte) error {
	var obj image.UploadImageRequest

	err := obj.UnmarshalJSON(msg)
	if err != nil {
		return err
	}

	if len(obj.Images) == 0 {
		return fmt.Errorf("no images provided")
	}

	for _, file := range obj.Images {
		file.Src, err = imageproc.CompressFile(file.Src, 75)
		if err != nil {
			return err
		}
	}

	return nil
}
