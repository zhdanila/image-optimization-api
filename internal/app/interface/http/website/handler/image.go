package handler

import (
	"github.com/labstack/echo/v4"
	"image-optimization-api/internal/service/image"
)

type Image struct {
	imageService *image.Service
}

func NewAuth(imageService *image.Service) *Image {
	return &Image{
		imageService: imageService,
	}
}

func (s *Image) Register(group *echo.Group) {

}
