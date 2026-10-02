package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		fmt.Print("JJJJ")
		c.JSON(200, gin.H{
			"message": "Hello World5",
		})
	})
}
