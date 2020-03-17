package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-web/app/model"
	"golang-web/app/rule/user"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
)

// Register user example
type Register struct {
	Email           string `json:"email" example:"google@gmail.com"`
	Password        string `json:"password" validate:"min=8,eqfield=ConfirmPassword"`
	ConfirmPassword string `json:"confirm_password"`
	VerifyCode      string `json:"verify_code" validate:"len=6"`
}

// 注册用户时的参数验证器
func RegisterValidate(c *gin.Context) (Register, error) {
	// 参数验证
	validate := validator.New()

	// 注册自定义验证
	_ = validate.RegisterValidation("NameValid", user.NameValid)

	newUser := Register{}

	_ = c.ShouldBindJSON(&newUser)

	err := validate.Struct(&newUser)

	return newUser, err
}

// 注册用户
func (u *User) Register(ru Register) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(ru.Password), bcrypt.DefaultCost)
	encodePW := string(hash)
	fmt.Println(len(encodePW))
	u.Name = "default name"
	u.Email = ru.Email
	u.Password = encodePW
	u.Avatar = "default avatar"
	model.Db.Create(&u)
}
