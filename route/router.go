package route

import (
	"github.com/gin-gonic/gin"
	"golang-web/app/controller/version1"
	"golang-web/route/middleware"
	"golang-web/route/middleware/logger"
)

func InitRouter() *gin.Engine {
	// 新建一个没有任何默认中间件的路由
	r := gin.New()

	// 全局中间件
	// Logger 中间件将日志写入 gin.DefaultWriter，即使你将 GIN_MODE 设置为 release。
	// By default gin.DefaultWriter = os.Stdout
	r.Use(logger.SetUp())

	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	//r.Use(gin.Recovery())

	// 你可以为每个路由添加任意数量的中间件。
	//r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	// 认证路由组
	apiV1 := r.Group("api/v1")

	// AuthRequired() 中间件
	apiV1.Use(middleware.AuthRequired())
	{
		//authorized.POST("/login", loginEndpoint)
		//authorized.POST("/submit", submitEndpoint)
		//authorized.POST("/read", readEndpoint)

		// 嵌套路由组
		apiV1.GET("/users", version1.Index)
		apiV1.GET("/users/:id", version1.Show)
		apiV1.POST("/users", version1.Store)
		apiV1.DELETE("/users/:id", version1.Destroy)
	}

	return r
}
