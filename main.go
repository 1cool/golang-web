package main

import (
	"golang-web/route"
)

func main() {
	r := route.InitRouter()

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
