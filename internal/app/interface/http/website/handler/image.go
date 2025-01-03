package handler

import (
	"github.com/labstack/echo/v4"
	"image-optimization-api/internal/service/image"
	"image-optimization-api/pkg/bind"
	"net/http"
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
	group.POST("/image", s.UploadImage)
}

func (s *Image) UploadImage(c echo.Context) error {
	var (
		err error
		obj image.UploadImageRequest
	)

	if err = bind.BindValidate(c, &obj); err != nil {
		return err
	}

	res, err := s.imageService.UploadImage(c.Request().Context(), &obj)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
