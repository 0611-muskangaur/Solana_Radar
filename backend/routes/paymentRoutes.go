// routes/paymentRoutes.go
package routes

import (
	"decentralised_payment_gateway/controllers"
	middleware "decentralised_payment_gateway/middlewares"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(r *gin.Engine) {
	payment := r.Group("/payments")
	payment.Use(middleware.AuthMiddleware())
	{
		payment.POST("/", controllers.CreatePaymentRequest)
		payment.GET("/", controllers.GetPayments)
	}
}
