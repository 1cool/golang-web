package rule

import (
	"golang-web/app/rule/user"
	"gopkg.in/go-playground/validator.v9"
)

func SetUp() {

	// 参数验证
	validate := validator.New()

	// 注册自定义验证
	_ = validate.RegisterValidation("NameValid", user.NameValid)

}
