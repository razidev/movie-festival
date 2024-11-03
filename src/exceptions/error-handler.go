package exception

import (
	"github.com/go-playground/validator/v10"
)

func ValidationError(err interface{}) map[string]string {
	validationErrors := err.(validator.ValidationErrors)
	errorsMap := make(map[string]string)
	for _, err := range validationErrors {
		fieldName := err.Field()
		switch err.Tag() {
		case "required":
			errorsMap[fieldName] = fieldName + " is required"
		case "min":
			errorsMap[fieldName] = fieldName + " must be at least " + err.Param() + " characters"
		case "max":
			errorsMap[fieldName] = fieldName + " must be at most " + err.Param() + " characters"
		case "email":
			errorsMap[fieldName] = "Invalid email format"
		case "http_url":
			errorsMap[fieldName] = "Movie Url is not valid"
		}
	}

	return errorsMap
}
