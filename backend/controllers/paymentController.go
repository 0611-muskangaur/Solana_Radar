package controllers

import (
	"decentralised_payment_gateway/models"
	"decentralised_payment_gateway/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid" // Import the UUID package
	"net/http"
	"strconv"
)

type PaymentController struct {
	service *services.PaymentService
}

func NewPaymentController(service *services.PaymentService) *PaymentController {
	return &PaymentController{service: service}
}

func (pc *PaymentController) CreatePaymentRequest(c *gin.Context) {
	var payment models.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	merchantIDFloat, exists := c.Get("merchant_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Merchant ID not found in context"})
		return
	}

	// Convert merchantID from float64 to uint
	merchantID := uint(merchantIDFloat.(float64))
	payment.MerchantID = merchantID

	// Generate a unique transaction hash
	payment.TransactionHash = uuid.New().String() // Generate a new UUID

	if err := pc.service.CreatePayment(&payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment request created", "payment_id": payment.ID})
}

func (pc *PaymentController) GetPayments(c *gin.Context) {
	merchantIDStr := c.Query("merchant_id") // Get merchant ID from query parameters

	// Check if the merchant ID is empty
	if merchantIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Merchant ID is required"})
		return
	}

	// Convert merchantIDStr to a float64 and handle potential errors
	merchantIDFloat, err := strconv.ParseFloat(merchantIDStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid merchant ID"})
		return
	}
	merchantID := uint(merchantIDFloat)

	// Get payments for the merchant
	payments, err := pc.service.GetPaymentsForMerchant(merchantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the payments in JSON format
	c.JSON(http.StatusOK, payments)
}

// HandleWebhook processes the webhook to update transaction status.
func (pc *PaymentController) HandleWebhook(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pc.service.UpdateTransactionStatus(&transaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction updated", "transaction_id": transaction.ID})
}
