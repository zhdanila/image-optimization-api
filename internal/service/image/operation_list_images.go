package image

import (
	"context"
	"image-optimization-api/internal/domain/image"
)

type operationListImages struct {
	*Service
	obj  *ListImageRequest
	ents []image.ImageInfo
}

func newOperationListImages(s *Service, obj *ListImageRequest) *operationListImages {
	return &operationListImages{
		Service: s,
		obj:     obj,
	}
}

func (o *operationListImages) listImages(ctx context.Context) error {
	var err error

	if o.ents, err = o.imageRepo.ListImages(ctx); err != nil {
		return err
	}

	return err
}

func (o *operationListImages) respond() *ListImageResponse {
	res := &ListImageResponse{
		Images: make([]ImageInfo, 0),
	}

	for _, ent := range o.ents {
		res.Images = append(res.Images, ImageInfo{
			Key: ent.Key,
			URL: ent.URL,
		})
	}

	return res
}
