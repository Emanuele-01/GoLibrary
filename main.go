package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"library.net/module/lib"
	"library.net/module/route"
	"library.net/module/services"
)

func main() {
	r := gin.Default()

	_, err := services.ConnectServiceMongo(lib.UriMongo, lib.DatabaseName)
	if err != nil {
		fmt.Println("Connection Mongo error: ", err.Error())
	}

	route.SetupRoute(r)

	r.Run(":8080")
}
