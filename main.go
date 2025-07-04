package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"library.net/module/config"
	"library.net/module/lib"
	"library.net/module/route"
)

func main() {
	r := gin.Default()

	mongo := config.NewDatabase(lib.DatabaseName, lib.UriMongo)

	errMGD := mongo.ConnectMDB(10, 50, 10)
	if errMGD != nil {
		log.Fatal("error connect database: ", errMGD.Error())
	}

	route.SetupRoute(r)

	log.Fatal(r.Run(fmt.Sprintf(":%d", config.Port)))
}
