package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// initialize router using default settings
	router := gin.Default()

	// initialize routes
	initializeRoutes(router)

	// start server
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
