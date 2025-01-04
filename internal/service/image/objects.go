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

type GetImageRequest struct {
	ImageID            string `param:"image_id" validate:"required"`
	CompressionQuality int    `query:"quality" validate:"omitempty,oneof=100 75 50 25"`
}

type GetImageResponse struct {
	Image ImageInfo `json:"image"`
}

type ListImageRequest struct{}

type ListImageResponse struct {
	Images []ImageInfo `json:"images"`
}

type ImageInfo struct {
	Key string `json:"key"`
	URL string `json:"url"`
}
