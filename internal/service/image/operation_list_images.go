package image

import (
	"context"
)

type operationListImages struct {
	*Service
	obj *ListImageRequest
}

func newOperationListImages(s *Service, obj *ListImageRequest) *operationListImages {
	return &operationListImages{
		Service: s,
		obj:     obj,
	}
}

func (o *operationListImages) listImages(ctx context.Context) error {
	var err error

	if err = o.imageRepo.ListImages(ctx); err != nil {
		return err
	}

	return err
}

func (o *operationListImages) respond() *ListImageResponse {
	res := &ListImageResponse{}

	return res
}
