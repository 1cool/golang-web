package version1

import (
	"github.com/gin-gonic/gin"
	"golang-web/app/model"
	"golang-web/app/rule"
	"golang-web/app/util/response"
	"gopkg.in/go-playground/validator.v9"
	"strconv"
)

// user 列表
func Index(c *gin.Context) {
	appG := response.Gin{C: c}
	user := model.User{}
	users := user.List()
	appG.Response(true, "", &users)
}

// 查看user详情
func Show(c *gin.Context) {
	appG := response.Gin{C: c}
	user := model.User{}
	users := user.List()

	intId, _ := strconv.Atoi(c.Param("id"))
	for _, user := range users {
		if user.Id == intId {
			appG.Response(true, "", &user)
			return
		}
	}

	appG.Response(true, "暂无数据", make([]string, 0))
}

// 添加 user
func Store(c *gin.Context) {
	appG := response.Gin{C: c}
	user := model.User{}
	users := user.List()

	// 参数验证
	validate := validator.New()

	// 注册自定义验证
	_ = validate.RegisterValidation("NameValid", rule.NameValid)

	newUser := model.User{}

	_ = c.ShouldBindJSON(&newUser)

	if err := validate.Struct(&newUser); err != nil {
		appG.Response(false, err.Error(), nil)
		return
	}

	gender, _ := strconv.Atoi(c.PostForm("gender"))

	id := len(users) + 1

	newUser.Id = id
	newUser.Gender = gender

	users = append(users, newUser)
	appG.Response(true, "", &users)
}

// 删除 user
func Destroy(c *gin.Context) {
	appG := response.Gin{C: c}
	user := model.User{}
	users := user.List()

	intId, _ := strconv.Atoi(c.Param("id"))

	for index, user := range users {
		if user.Id == intId {
			users = append(users[:index], users[index+1:]...)
		}
	}

	appG.Response(true, "", &users)
}
