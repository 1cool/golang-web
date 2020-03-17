package user

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8;"`
}

func LoginValidate(c *gin.Context) error {
	// 参数验证
	validate := validator.New()

	newUser := Register{}

	_ = c.ShouldBindJSON(&newUser)

	err := validate.Struct(&newUser)

	return err
}
