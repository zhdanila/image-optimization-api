package image

import (
	"image-optimization-api/pkg/bind"
)

const imagesLimit = 100

type UploadImageRequest struct {
	Images []bind.UploadedFile `form:"images" validate:"omitempty,dive"`
}

func (r *UploadImageRequest) ImagesToFill() []bind.UploadedFile {
	if r.Images == nil {
		r.Images = make([]bind.UploadedFile, imagesLimit)
		return r.Images
	}

	var filledImages []bind.UploadedFile
	for _, img := range r.Images {
		if img.Size > 0 {
			filledImages = append(filledImages, img)
		}
	}

	return filledImages
}

type UploadImageResponse struct{}

type GetImageRequest struct {
	ListingId string `query:"image_id" validate:"required"`
}

type GetImageResponse struct{}

type ListImageRequest struct{}

type ListImageResponse struct{}
