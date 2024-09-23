// controllers/paymentController.go
package controllers

import (
	"decentralised_payment_gateway/models"
	"decentralised_payment_gateway/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePaymentRequest(c *gin.Context) {
	var payment models.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	services.CreatePayment(&payment)
	c.JSON(http.StatusOK, gin.H{"message": "Payment request created"})
}

func GetPayments(c *gin.Context) {
	merchantID, _ := c.Get("merchant_id") // Extract merchant ID from JWT
	payments := services.GetPaymentsForMerchant(merchantID.(uint))
	c.JSON(http.StatusOK, payments)
}
