package handlers

import (
	"e-commerce/internal/utils"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func init() {
	// validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
	// 	password := fl.Field().String()
	// 	re := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&]).{8,}$`)
	// 	return re.MatchString(password)
	// })
	validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
	return utils.IsStrongPassword(fl.Field().String())
})
}
