package user

import (
	"gopkg.in/go-playground/validator.v9"
)

func NameValid(fl validator.FieldLevel) bool {
	name := fl.Field().String()
	if len(name) < 2 {
		return false
	}

	return true
}
