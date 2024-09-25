// main.go
package main

import (
	"decentralised_payment_gateway/config"
	"decentralised_payment_gateway/db"
	"decentralised_payment_gateway/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	db.ConnectPostgres()
	db.AutoMigrate() // Automatically create tables
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Define a handler for the root path
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to the Decentralized Payment Gateway!")
	})

	// Set up your application routes
	routes.MerchantRoutes(r)
	routes.PaymentRoutes(r)

	// Start the server
	r.Run(":8080")
}
