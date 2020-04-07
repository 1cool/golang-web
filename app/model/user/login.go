package user

import (
	"github.com/gin-gonic/gin"
	"golang-web/app/model"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
)

type LoginParam struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func LoginValidate(c *gin.Context) (LoginParam, error) {
	// 参数验证
	validate := validator.New()

	var login LoginParam

	_ = c.ShouldBindJSON(&login)

	err := validate.Struct(&login)

	return login, err
}

func Login(lu LoginParam) (User, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(lu.Password), bcrypt.DefaultCost)

	encodePW := string(hash)
	var user User
	_ = model.Db.Where("email = ?", lu.Email).First(&user)

	// 正确密码验证
	err := bcrypt.CompareHashAndPassword([]byte(encodePW), []byte(user.Password))

	return user, err
}
