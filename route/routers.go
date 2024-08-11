package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoute(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusAccepted, gin.H{"message": "Pong"})
	})

	//books route
	book := r.Group("book")
	{
		book.GET("/:id")
		book.GET("/all")
		book.POST("/create")
		book.PUT("/mod")
		book.DELETE("/del")
	}

	//admin route
	admin := r.Group("admin")
	{
		admin.GET("/:id")
		admin.GET("/all")
		admin.POST("/create")
		admin.PUT("/mod")
		admin.DELETE("/del")
	}

	//user route
	users := r.Group("user")
	{
		users.GET("/:id")
		users.GET("/all")
		users.POST("/create")
		users.PUT("/mod")
		users.DELETE("/del")
	}
}
