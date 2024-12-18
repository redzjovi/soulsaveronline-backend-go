package http

import "github.com/go-playground/validator/v10"

func NewMapErrorJson(err error) map[string]string {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		mapError := make(map[string]string)

		for _, e := range validationErrors {

			mapError[e.Field()] = e.Tag()
		}

		return mapError
	}

	return map[string]string{"error": "Invalid input"}
}
