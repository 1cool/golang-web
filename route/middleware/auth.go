package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"golang-web/app/util/response"
	"golang-web/config"
	"log"
	"net/http"
	"strconv"
	"time"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		var AppConfig config.App

		_ = envconfig.Process("myapp", &AppConfig)

		tsGap := AppConfig.TsGap

		tsInt, _ := strconv.ParseInt(c.GetHeader("ts"), 10, 64)

		if tsInt == 0 || tsInt > t.Unix() || t.Unix()-tsInt > tsGap {
			appG := response.Gin{C: c}
			appG.Response(http.StatusOK, false, "请求无效", []string{})
			c.Abort()
			return
		}

		// 设置 example 变量
		c.Set("example", "12345")

		// 请求前

		c.Next()

		// 请求后
		latency := time.Since(t)
		log.Print(latency)

		// 获取发送的 status
		status := c.Writer.Status()
		log.Println(status)
	}
}
