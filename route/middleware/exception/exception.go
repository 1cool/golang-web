package exception

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"golang-web/app/util/mail"
	"golang-web/app/util/response"
	"golang-web/app/util/time"
	"golang-web/config"
	"runtime/debug"
	"strings"
)

func SetUp() gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				DebugStack := ""

				for _, v := range strings.Split(string(debug.Stack()), "\n") {
					DebugStack += v + "<br>"
				}

				var AppConfig config.App
				_ = envconfig.Process("myapp", &AppConfig)
				subject := fmt.Sprintf("【重要错误】%s 项目出错了！", AppConfig.AppName)

				body := strings.ReplaceAll(MailTemplate, "{ErrorMsg}", fmt.Sprintf("%s", err))
				body = strings.ReplaceAll(body, "{RequestTime}", time.GetCurrentTimeByFormat("2006-01-02 15:04:05"))
				body = strings.ReplaceAll(body, "{RequestURL}", c.Request.Method+"  "+c.Request.Host+c.Request.RequestURI)
				body = strings.ReplaceAll(body, "{RequestUA}", c.Request.UserAgent())
				body = strings.ReplaceAll(body, "{RequestIP}", c.ClientIP())
				body = strings.ReplaceAll(body, "{DebugStack}", DebugStack)

				mailOption := &mail.Option{
					Host:     "",
					Port:     0,
					User:     "",
					Password: "",
					To:       "",
					Subject:  subject,
					Body:     "",
				}

				mail.Send(mailOption)

				appG := response.Gin{C: c}
				appG.Response(500, false, DebugStack, []string{})
			}
		}()

		c.Next()
	}
}
