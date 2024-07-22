package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoute(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusAccepted, gin.H{"message": "Pong"})
	})
}
