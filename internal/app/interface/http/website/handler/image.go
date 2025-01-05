package handler

import (
	"github.com/labstack/echo/v4"
	image2 "image-optimization-api/internal/domain/image"
	"image-optimization-api/internal/service/image"
	"image-optimization-api/pkg/bind"
	_ "image-optimization-api/pkg/server"
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
	group.GET("/origin", s.ListOriginImages)
}

// UploadImage @Summary Upload Image
// @Description Uploads an image
// @Tags Images
// @ID upload-image
// @Produce  json
// @Param images formData file true "Images to upload"
// @Success 200 {object} server.EmptyResponse
// @Router /api/image [post]
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

// GetImage @Summary Get Image
// @Description Retrieves an image by ID
// @Tags Images
// @ID get-image
// @Produce  json
// @Param image_id path string true "Image ID"
// @Param quality query int false "Compression Quality (one of: 100, 75, 50, 25)"
// @Success 200 {object} image.GetImageResponse
// @Router /api/image/{image_id} [get]
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

// ListImages @Summary List Images
// @Description Retrieves a list of images
// @Tags Images
// @ID list-images
// @Produce  json
// @Success 200 {object} image.ListImageResponse
// @Router /api/image/list [get]
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

// ListOriginImages @Summary List Original Images
// @Description Retrieves a list of original images, excluding those with compression quality suffixes
// @Tags Images
// @ID list-origin-images
// @Produce json
// @Success 200 {object} image.ListOriginImageResponse
// @Router /api/image/origin [get]
func (s *Image) ListOriginImages(c echo.Context) error {
	var (
		err error
		obj image.ListOriginImageRequest
	)

	if err = bind.BindValidate(c, &obj); err != nil {
		return err
	}

	res, err := s.imageService.ListOriginImages(c.Request().Context(), &obj)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
