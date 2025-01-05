package image

import (
	"context"
	"image-optimization-api/internal/domain/image"
	"image-optimization-api/pkg/imageproc"
)

type operationGetImage struct {
	*Service
	imageID string
	obj     *GetImageRequest
	ent     *image.Info
}

func newOperationGetImage(s *Service, obj *GetImageRequest) *operationGetImage {
	return &operationGetImage{
		Service: s,
		obj:     obj,
	}
}

func (o *operationGetImage) prepareImageID(_ context.Context) error {
	var err error

	quality := imageproc.IntToCompressionQuality(o.obj.CompressionQuality)
	o.imageID = imageproc.GenerateImageID(o.obj.ImageID, quality)

	return err
}

func (o *operationGetImage) getImage(ctx context.Context) error {
	var err error

	if o.ent, err = o.imageRepo.GetImage(ctx, o.imageID); err != nil {
		return errs.ImageNotFound.SetInternal(err)
	}

	return err
}

func (o *operationGetImage) respond() *GetImageResponse {
	res := &GetImageResponse{
		Info{
			Key: o.ent.Key,
			URL: o.ent.URL,
		},
	}

	return res
}
