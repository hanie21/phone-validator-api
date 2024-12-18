package main

import (
	"PhoneValidatorAPI/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Define the route and associate it with the handler function
	r.GET("/v1/phone-numbers", handlers.PhoneNumberHandler)

	// Define the route and associate it with the POST handler function
	r.POST("/v1/phone-numbers", handlers.CreatePhoneNumberHandler)

	// Start the server
	r.Run(":8080")
}
