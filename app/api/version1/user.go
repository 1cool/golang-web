package version1

import (
	"github.com/gin-gonic/gin"
	"golang-web/app/model"
	"golang-web/app/rule"
	"golang-web/app/util/response"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

// @Tags 用户相关
// @Summary 用户列表
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/users [get]
func Index(c *gin.Context) {
	appG := response.Gin{C: c}
	appG.Response(http.StatusOK, true, "暂无数据", make([]string, 0))
}

// @Tags 用户相关
// @Summary 用户详情
// @Produce  json
// @Param id path int true "用户id"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/users/{id} [get]
func Show(c *gin.Context) {
	appG := response.Gin{C: c}

	appG.Response(http.StatusOK, true, "暂无数据", make([]string, 0))
}

// 添加 user
// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v1
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
