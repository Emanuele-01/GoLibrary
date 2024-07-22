package main

import (
	"github.com/gin-gonic/gin"
	"library.net/module/route"
)

func main() {
	r := gin.Default()

	route.SetupRoute(r)

	r.Run(":8080")
}
