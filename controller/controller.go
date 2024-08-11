package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"library.net/module/services"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		fmt.Println("Nessun Parametr utente trovao all'inetrno della stringa")
		c.JSON(http.StatusNotFound, gin.H{
			"message": "nessun utente trovato",
		})
		return
	}

	db, err := services.UserDatabaseConnect()
	if err != nil {
		fmt.Println("Error Controller Line 24: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, err := db.GetUser(id)
	if err != nil {
		fmt.Println("Error Controller Line 33: ", err.Error())
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
