package image

import (
	"context"
	"image-optimization-api/pkg/rabbitmq"
)

type operationQueuePublish struct {
	*Service
	obj *UploadImageRequest
}

func newOperationQueuePublish(s *Service, obj *UploadImageRequest) *operationQueuePublish {
	return &operationQueuePublish{
		Service: s,
		obj:     obj,
	}
}

func (o *operationQueuePublish) queuePublish(ctx context.Context) error {
	var err error

	body, err := o.obj.MarshalJSON()
	if err != nil {
		return err
	}

	if err = rabbitmq.PublishToQueue(o.conn, body); err != nil {
		return err
	}

	return nil
}

func (o *operationQueuePublish) respond() *UploadImageResponse {
	res := &UploadImageResponse{}

	return res
}
