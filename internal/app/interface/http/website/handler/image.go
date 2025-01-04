package handler

import (
	"github.com/labstack/echo/v4"
	image2 "image-optimization-api/internal/domain/image"
	"image-optimization-api/internal/service/image"
	"image-optimization-api/pkg/bind"
	"net/http"
)

type Image struct {
	imageService *image.Service
}

func NewImage(imageService *image.Service) *Image {
	return &Image{
		imageService: imageService,
	}
}

func (s *Image) Register(server *echo.Group) {
	group := server.Group("/image")

	group.POST("", s.UploadImage)
	group.GET("/:image_id", s.GetImage)
	group.GET("/list", s.ListImages)
}

func (s *Image) UploadImage(c echo.Context) error {
	var (
		err error
		obj image.UploadImageRequest
	)

	if err = bind.BindValidate(c, &obj, bind.FromHeaders()); err != nil {
		return err
	}

	for _, group := range image2.Groups() {
		if err = bind.BindValidate(c, &obj, bind.FromMultipartForm(group, obj.ImagesToFill())); err != nil {
			return err
		}
	}

	res, err := s.imageService.UploadImage(c.Request().Context(), &obj)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (s *Image) GetImage(c echo.Context) error {
	var (
		err error
		obj image.GetImageRequest
	)

	if err = bind.BindValidate(c, &obj); err != nil {
		return err
	}

	res, err := s.imageService.GetImage(c.Request().Context(), &obj)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (s *Image) ListImages(c echo.Context) error {
	var (
		err error
		obj image.ListImageRequest
	)

	if err = bind.BindValidate(c, &obj); err != nil {
		return err
	}

	res, err := s.imageService.ListImages(c.Request().Context(), &obj)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
