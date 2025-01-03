package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

func NewValidator() *Validator {
	return &Validator{validator.New()}
}

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) RegisterDefaultRules() error {
	var err error

	if err = v.validator.RegisterValidation("uuid", ValidateUUID); err != nil {
		return err
	}
	if err = v.validator.RegisterValidation("timestamp", ValidateTimestamp); err != nil {
		return err
	}

	return nil
}

func (v *Validator) Validate(i any) error {
	return v.validator.Struct(i)
}

func (v *Validator) Register(tag string, fn validator.Func) error {
	return v.validator.RegisterValidation(tag, fn)
}

// ValidateUUID is a custom validation function for UUIDs.
func ValidateUUID(fl validator.FieldLevel) bool {
	_, err := uuid.Parse(fl.Field().String())
	return err == nil
}

// ValidateTimestamp is a custom validation function for timestamps.
func ValidateTimestamp(fl validator.FieldLevel) bool {
	_, err := time.Parse(time.RFC3339, fl.Field().String())
	return err == nil
}
