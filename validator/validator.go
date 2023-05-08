package validator

import (
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func New() *Validator {
	v := validator.New()
	return &Validator{
		validator: v,
	}
}

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) ValidateSimple(i interface{}) error {
	return v.validator.Struct(i)
}

func (v *Validator) Validate(i interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := v.validator.Struct(i)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
