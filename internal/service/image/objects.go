package image

import "mime/multipart"

type UploadImageRequest struct {
	Image *multipart.FileHeader `form:"image" validate:"required"`
}

type UploadImageResponse struct{}

type GetImageRequest struct {
	ListingId string `query:"image_id" validate:"required"`
}

type GetImageResponse struct{}

type ListImageRequest struct{}

type ListImageResponse struct{}
