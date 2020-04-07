package v1

import (
	"github.com/gin-gonic/gin"
	"golang-web/app/model/user"
	"golang-web/app/util/response"
	"net/http"
)

// @Tags 用户相关
// @Summary 用户登录
// @Produce json
// @Accept  json
// @Param user body user.LoginParam true "用户登录"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/users/auth [POST]
func Login(c *gin.Context) {
	appG := response.Gin{C: c}

	lu, err := user.LoginValidate(c)

	if err != nil {
		appG.Response(http.StatusOK, true, err.Error(), make([]string, 0))
		return
	}

	var u user.User
	u, err = user.Login(lu)

	if err != nil {
		appG.Response(http.StatusOK, true, "账号或者密码错误", make([]string, 0))
		return
	}

	appG.Response(http.StatusOK, true, "", u)
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
// @Param user body user.Register true "用户注册"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/users [post]
func Register(c *gin.Context) {
	appG := response.Gin{C: c}

	newUser, err := user.RegisterValidate(c)

	if err != nil {
		appG.Response(http.StatusOK, true, err.Error(), make([]string, 0))
		return
	}

	u := user.User{}
	u.Register(newUser)

	appG.Response(http.StatusOK, true, "暂无数据", make([]string, 0))
}

// 删除 user
func Destroy(c *gin.Context) {
	appG := response.Gin{C: c}

	appG.Response(http.StatusOK, true, "暂无数据", make([]string, 0))
}
