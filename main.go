package main

import (
	"golang-web/app/model"
	"golang-web/config"
	"golang-web/route"
	"golang-web/route/middleware/exception"
)

func init() {
	config.Setup()
	exception.SetUp()
	model.SetUp()
}

// @title github.com/1cool/golang-web Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService https://github.com/1cool/golang-web

// @contact.name 1cool
// @contact.url https://github.com/1cool/golang-web
// @contact.email 1coolluobo@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	r := route.InitRouter()

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
