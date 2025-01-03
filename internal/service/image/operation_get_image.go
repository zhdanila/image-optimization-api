package image

import (
	"context"
)

type operationGetImage struct {
	*Service
	obj *GetImageRequest
}

func newOperationGetImage(s *Service, obj *GetImageRequest) *operationGetImage {
	return &operationGetImage{
		Service: s,
		obj:     obj,
	}
}

func (o *operationGetImage) getImage(ctx context.Context) error {
	var err error

	if err = o.imageRepo.GetImage(ctx); err != nil {
		return err
	}

	return err
}

func (o *operationGetImage) respond() *GetImageResponse {
	res := &GetImageResponse{}

	return res
}
