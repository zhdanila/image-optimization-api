package compression

import (
	"context"
	"image-optimization-api/internal/service/image"
	"image-optimization-api/pkg/bind"
	"image-optimization-api/pkg/imageproc"
)

type operationCompressImages struct {
	*Service
	body             []byte
	obj              *image.UploadImageRequest
	compressedImages []bind.UploadedFile
}

func newOperationCompressImages(service *Service, body []byte) *operationCompressImages {
	return &operationCompressImages{Service: service, body: body}
}

func (o *operationCompressImages) unmarshalBody(ctx context.Context) error {
	var obj image.UploadImageRequest

	err := obj.UnmarshalJSON(o.body)
	if err != nil {
		return err
	}

	if obj.Images == nil || len(obj.Images) == 0 {
		return errs.NoImagesProvided
	}

	o.obj = &obj

	return nil
}

func (o *operationCompressImages) compressImage(ctx context.Context) error {
	var err error

	o.compressedImages, err = imageproc.GetCompressedImages(o.obj.Images)
	if err != nil {
		return err
	}

	return nil
}

func (o *operationCompressImages) uploadImages(ctx context.Context) error {
	var err error

	if err = o.imageRepo.UploadImages(ctx, o.compressedImages); err != nil {
		return err
	}

	return nil
}
