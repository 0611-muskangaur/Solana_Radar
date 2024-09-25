package routes

import (
	"decentralised_payment_gateway/controllers"
	middleware "decentralised_payment_gateway/middlewares"

	"github.com/gin-gonic/gin"
)

// PaymentRoutes sets up the routes for payment-related operations.
func PaymentRoutes(r *gin.Engine, paymentController *controllers.PaymentController) {
	payment := r.Group("/payments")
	payment.Use(middleware.AuthMiddleware()) // Apply authentication middleware
	{
		payment.POST("/", paymentController.CreatePaymentRequest) // Endpoint to create a payment
		payment.GET("/", paymentController.GetPayments)           // Endpoint to get payments for a merchant
		payment.POST("/webhook", paymentController.HandleWebhook) // Webhook for transaction updates
	}
}
