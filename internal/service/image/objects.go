package image

type UploadImageRequest struct{}

type UploadImageResponse struct{}

type GetImageRequest struct {
	ListingId string `query:"image_id" validate:"required"`
}

type GetImageResponse struct{}
