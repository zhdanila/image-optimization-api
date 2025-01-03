package image

import (
	"context"
)

type operationUploadImage struct {
	*Service
	obj *UploadImageRequest
}

func newOperationUploadImage(s *Service, obj *UploadImageRequest) *operationUploadImage {
	return &operationUploadImage{
		Service: s,
		obj:     obj,
	}
}

func (o *operationUploadImage) uploadImage(ctx context.Context) error {
	var err error

	if err = o.imageRepo.UploadImage(ctx); err != nil {
		return err
	}

	return err
}

func (o *operationUploadImage) respond() *UploadImageResponse {
	res := &UploadImageResponse{}

	return res
}
