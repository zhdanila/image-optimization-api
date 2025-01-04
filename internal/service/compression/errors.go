package compression

import (
	rest "image-optimization-api/pkg/serializer"
	"net/http"
)

var errs = struct {
	NoImagesProvided *rest.Error
}{
	NoImagesProvided: rest.NewError(http.StatusBadRequest, "No images provided"),
}
