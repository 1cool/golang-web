package main

import (
	"golang-web/routes"
)

func main() {
	//r := gin.Default()

	r := routes.InitRouter()
	//r.GET("/ping", func(c *gin.Context) {
	//	data := map[string]interface{}{
	//		"lang": "GO语言",
	//		"tag":  "<br>",
	//	}
	//	c.JSON(http.StatusOK, data)
	//})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
