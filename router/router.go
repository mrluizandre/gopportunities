package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// initialize router using default settings
	router := gin.Default()

	// define a route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// start server
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
