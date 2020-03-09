package main

import (
	"golang-web/config"
	"golang-web/route"
	"golang-web/route/middleware/exception"
)

func init() {
	config.Setup()
	exception.SetUp()
}

func main() {
	r := route.InitRouter()

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
