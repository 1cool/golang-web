package version1

import (
	"github.com/gin-gonic/gin"
	"golang-web/app/model"
	"golang-web/app/rule"
	"golang-web/app/util/response"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

// user 列表
func Index(c *gin.Context) {
	appG := response.Gin{C: c}
	appG.Response(http.StatusOK, true, "暂无数据", make([]string, 0))
}

// 查看user详情
func Show(c *gin.Context) {
	appG := response.Gin{C: c}

	appG.Response(http.StatusOK, true, "暂无数据", make([]string, 0))
}

// 添加 user
func Store(c *gin.Context) {
	appG := response.Gin{C: c}

	// 参数验证
	validate := validator.New()

	// 注册自定义验证
	_ = validate.RegisterValidation("NameValid", rule.NameValid)

	newUser := model.User{}

	_ = c.ShouldBindJSON(&newUser)

	if err := validate.Struct(&newUser); err != nil {
		appG.Response(http.StatusUnprocessableEntity, false, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, true, "暂无数据", make([]string, 0))
}

// 删除 user
func Destroy(c *gin.Context) {
	appG := response.Gin{C: c}

	appG.Response(http.StatusOK, true, "暂无数据", make([]string, 0))
}
