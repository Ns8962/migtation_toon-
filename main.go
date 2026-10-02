// package main

// import (
//     "github.com/gin-gonic/gin"
//     "net/http"

// )

// func main() {
//     r := gin.Default()

//     // GET /ping
//     r.GET("/ping", func(c *gin.Context) {
//         c.JSON(http.StatusOK, gin.H{
//             "message": "pong",
//         })
//     })

//     // POST /hello
//     r.POST("/hello", func(c *gin.Context) {
//         var data struct {
//             Name string `json:"name"`
//         }
//         if err := c.ShouldBindJSON(&data); err != nil {
//             c.JSON(400, gin.H{"error": err.Error()})
//             return
//         }

//         c.JSON(200, gin.H{
//             "message": "Hello " + data.Name,
//         })
//     })

//	    r.Run(":8080")
//	}
package main

import (
	"service_go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
