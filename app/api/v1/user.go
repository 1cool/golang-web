package v1

import (
	"github.com/gin-gonic/gin"
	"golang-web/app/model"
	"golang-web/app/rule/user"
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

// @Tags 用户相关
// @Summary 用户注册
// @Produce  json
// @Accept  json
// @Param user body model.UserRegister true "用户注册"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/users [post]
func Store(c *gin.Context) {
	appG := response.Gin{C: c}

	// 参数验证
	validate := validator.New()

	// 注册自定义验证
	_ = validate.RegisterValidation("NameValid", user.NameValid)

	newUser := model.UserRegister{}

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
