package utils

import (
	"reflect"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func SetupValidator() {
	Validate = validator.New()
	Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("json")
		if name == "-" {
			return ""
		}
		return name
	})

	if err := Validate.RegisterValidation("password_validator", PasswordValidator); err != nil {
		return
	}
}

func PasswordValidator(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*()+-_=\[\]{}|;:,.<>?]`).MatchString(password)

	return hasUpper && hasNumber && hasSpecial && hasLower
}
func ValidateStruct(payload interface{}) error {
	Validate.SetTagName("binding")
	return Validate.Struct(payload)
}
