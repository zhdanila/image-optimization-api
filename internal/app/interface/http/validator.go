package http

import (
	"github.com/go-playground/validator/v10"
)

func CustomValidator() *Validator {
	return NewValidator()
}

func NewValidator() *Validator {
	return &Validator{validator: validator.New()}
}

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i any) error {
	return v.validator.Struct(i)
}
