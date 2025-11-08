package main

import "RoutePilot/api"

// import "github.com/gin-gonic/gin"

func main() {

	// api.LoginHandler(c * gin.Context)
	// Initialize the router
	// router := gin.Default()
	api.ApiServer()
	// // // Register the public API routes
	// router.POST("/login", LoginHandler)
	// router.POST("/register", RegisterHandler)

	// // Start the server
	// router.Run(":8080")
}
