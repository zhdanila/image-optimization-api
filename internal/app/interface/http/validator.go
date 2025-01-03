package http

import (
	"image-optimization-api/pkg/validator"
)

func CustomValidator() *validator.Validator {
	return validator.NewValidator()
}
