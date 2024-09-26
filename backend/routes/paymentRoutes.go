package routes

import (
	"decentralised_payment_gateway/controllers"
	"decentralised_payment_gateway/middlewares"

	"github.com/gin-gonic/gin"
)

// PaymentRoutes sets up the payment-related endpoints.
func PaymentRoutes(r *gin.Engine, paymentController *controllers.PaymentController) {
	payment := r.Group("/payments")
	payment.Use(middleware.AuthMiddleware())
	{
		payment.POST("/", paymentController.CreatePaymentRequest)
		payment.GET("/", paymentController.GetPayments)
		payment.POST("/webhook", paymentController.HandleWebhook)
	}
}
