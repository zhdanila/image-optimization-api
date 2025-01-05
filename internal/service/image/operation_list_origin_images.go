package image

import (
	"context"
	"image-optimization-api/internal/domain/image"
	"image-optimization-api/pkg/imageproc"
	"strings"
)

type operationListOriginImages struct {
	*Service
	obj  *ListOriginImageRequest
	ents []image.Info
	res  []string
}

func newOperationListOriginImages(s *Service, obj *ListOriginImageRequest) *operationListOriginImages {
	return &operationListOriginImages{
		Service: s,
		obj:     obj,
	}
}

func (o *operationListOriginImages) listOriginImages(ctx context.Context) error {
	var err error

	if o.ents, err = o.imageRepo.ListImages(ctx); err != nil {
		return err
	}

	return err
}

func (o *operationListOriginImages) findOriginalImage(_ context.Context) error {
	prefixes := imageproc.GetCompressionQualitySuffix()

	for _, ent := range o.ents {
		isForbidden := false
		for _, prefix := range prefixes {
			if strings.HasPrefix(ent.Key, prefix) {
				isForbidden = true
				break
			}
		}

		if !isForbidden {
			o.res = append(o.res, ent.Key)
		}
	}

	return nil
}

func (o *operationListOriginImages) respond() *ListOriginImageResponse {
	return &ListOriginImageResponse{
		Keys: o.res,
	}
}
