package utils

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

func ValidationErrors(err error) map[string]string {
	errors := map[string]string{}

	if err == nil {
		return errors
	}

	for _, err := range err.(validator.ValidationErrors) {
		field := err.Field()
		field = toSnakeCase(field)

		switch err.Tag() {
		case "required":
			errors[field] = field + " is required"
		case "email":
			errors[field] = "invalid email format"
		case "min":
			errors[field] = field + " must be at least " + err.Param() + " characters"
		case "max":
			errors[field] = field + " must be at most " + err.Param() + " characters"
		case "uuid":
			errors[field] = field + " must be a valid UUID"
		default:
			errors[field] = "invalid value for " + field
		}
	}

	return errors
}

func toSnakeCase(s string) string {
	var output []rune
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}