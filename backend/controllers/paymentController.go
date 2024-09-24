// controllers/paymentController.go
package controllers

import (
	"decentralised_payment_gateway/models"
	"decentralised_payment_gateway/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePaymentRequest(c *gin.Context) {
	var payment models.Payment //Receives a JSON payload representing a payment.
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return //Returns an HTTP 200 OK status with a success message if the payment request is created, or a 400 Bad Request status if there's an issue with the input.
	}

	services.CreatePayment(&payment) //Retrieves the merchant_id from the request context, which is set by the JWT authentication middleware.
	c.JSON(http.StatusOK, gin.H{"message": "Payment request created"})
}

func GetPayments(c *gin.Context) {
	merchantID, _ := c.Get("merchant_id")                          // Extract merchant ID from JWT
	payments := services.GetPaymentsForMerchant(merchantID.(uint)) //Fetches all payments associated with the merchant using services.GetPaymentsForMerchant, passing the merchant_id.
	c.JSON(http.StatusOK, payments)
}
