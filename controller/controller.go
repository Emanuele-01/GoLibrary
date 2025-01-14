package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"library.net/module/config"
	"library.net/module/services"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		log.Println("Nessun Parametr utente trovao all'inetrno della stringa")
		c.JSON(http.StatusNotFound, gin.H{
			"message": "nessun utente trovato",
		})
		return
	}

	usrService := services.NewUsrService(config.App.D)

	user, err := usrService.GetUser(id)
	if err != nil {
		log.Println("Error Controller Line 33: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    user,
	})
}
