package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Custom validation error
type ValidationError struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Value   string `json:"value"`
	Message string `json:"message"`
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation failed on field '%s': %s", e.Field, e.Message)
}

// Validator wrapper
type Validator struct {
	validate *validator.Validate
}

// NewValidator creates a new validator instance
func NewValidator() *Validator {
	return &Validator{
		validate: validator.New(),
	}
}

// ValidateStruct validates a struct based on its tags
func (v *Validator) ValidateStruct(s interface{}) []ValidationError {
	var errors []ValidationError

	err := v.validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, ValidationError{
				Field:   err.Field(),
				Tag:     err.Tag(),
				Value:   fmt.Sprintf("%v", err.Value()),
				Message: getErrorMessage(err),
			})
		}
	}

	return errors
}

// getErrorMessage returns a user-friendly error message based on the validation tag
func getErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("Field '%s' is required", err.Field())
	case "min":
		return fmt.Sprintf("Field '%s' must be at least %s characters", err.Field(), err.Param())
	case "max":
		return fmt.Sprintf("Field '%s' must be at most %s characters", err.Field(), err.Param())
	case "email":
		return fmt.Sprintf("Field '%s' must be a valid email address", err.Field())
	case "len":
		return fmt.Sprintf("Field '%s' must be exactly %s characters", err.Field(), err.Param())
	// Account-specific validations
	case "oneof":
		return fmt.Sprintf("Field '%s' must be one of the following values: %s", err.Field(), err.Param())
	default:
		return fmt.Sprintf("Field '%s' failed validation", err.Field())
	}
}
