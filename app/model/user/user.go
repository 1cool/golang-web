package user

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"golang-web/app/rule/user"
	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100);unique_index;not null" json:"email" validate:"required,email"`
	Password string `gorm:"type:varchar(36);not null" json:"password" validate:"required,min=8"`
	Name     string `gorm:"size:20;not null" json:"name"`
	Avatar   string `gorm:"size:255;not null" json:"avatar"`
	Point    int64  `gorm:"size:255;default:0;not null" json:"point"`
}

// Register user example
type Register struct {
	Email           string `json:"email" example:"google@gmail.com"`
	Password        string `json:"password" validate:"eqfield=ConfirmPassword"`
	ConfirmPassword string `json:"confirm_password"`
	VerifyCode      string `json:"verify_code" validate:"len=6"`
}

func RegisterValidate(c *gin.Context) error {
	// 参数验证
	validate := validator.New()

	// 注册自定义验证
	_ = validate.RegisterValidation("NameValid", user.NameValid)

	newUser := Register{}

	_ = c.ShouldBindJSON(&newUser)

	err := validate.Struct(&newUser)

	return err
}
