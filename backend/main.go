package main

import (
	"decentralised_payment_gateway/config"
	"decentralised_payment_gateway/controllers"
	"decentralised_payment_gateway/db"
	"decentralised_payment_gateway/routes"
	"decentralised_payment_gateway/services"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Connect to PostgreSQL database
	db.ConnectPostgres()

	// Automatically create tables based on models
	db.AutoMigrate()

	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)

	// Create a new Gin router
	r := gin.Default()

	// Define a handler for the root path
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to the Decentralized Payment Gateway!")
	})

	// Initialize services and controllers
	paymentService := services.NewPaymentService(db.GetDB()) // Assuming db.GetDB() returns a *gorm.DB
	paymentController := controllers.NewPaymentController(paymentService)

	// Set up merchant routes
	routes.MerchantRoutes(r)

	// Set up payment routes
	routes.PaymentRoutes(r, paymentController)

	// Start the server
	r.Run(":8080")
}
