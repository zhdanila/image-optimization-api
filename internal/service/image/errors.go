package image

import (
	rest "image-optimization-api/pkg/serializer"
	"net/http"
)

var errs = struct {
	ImageNotFound *rest.Error
}{
	ImageNotFound: rest.NewError(http.StatusBadRequest, "Image not found"),
}
