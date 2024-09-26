package controllers

import (
	"decentralised_payment_gateway/models"
	"decentralised_payment_gateway/res"
	"decentralised_payment_gateway/services"
	"fmt" // Make sure to import fmt for logging
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strconv"
)

type PaymentController struct {
	service *services.PaymentService
}

func NewPaymentController(service *services.PaymentService) *PaymentController {
	return &PaymentController{service: service}
}

// CreatePaymentRequest handles the creation of a new payment request.
func (pc *PaymentController) CreatePaymentRequest(c *gin.Context) {
	var payment models.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		res.BadRequest(c, err)
		return
	}

	merchantIDFloat, exists := c.Get("merchant_id")
	if !exists {
		res.Unauthorized(c, nil)
		return
	}

	merchantID := uint(merchantIDFloat.(float64))
	payment.MerchantID = merchantID
	payment.TransactionHash = uuid.New().String()

	if err := pc.service.CreatePayment(&payment); err != nil {
		res.InternalServerError(c, err)
		return
	}

	res.SuccessOK(c, gin.H{"message": "Payment request created", "payment_id": payment.ID})
}

// GetPayments fetches payment requests for a specific merchant.
func (pc *PaymentController) GetPayments(c *gin.Context) {
	merchantIDStr := c.Query("merchant_id")
	if merchantIDStr == "" {
		res.BadRequest(c, nil)
		return
	}

	merchantIDFloat, err := strconv.ParseFloat(merchantIDStr, 64)
	if err != nil {
		res.BadRequest(c, err)
		return
	}
	merchantID := uint(merchantIDFloat)

	payments, err := pc.service.GetPaymentsForMerchant(merchantID)
	if err != nil {
		res.InternalServerError(c, err)
		return
	}

	res.SuccessOK(c, payments)
}

// HandleWebhook handles transaction status updates via webhook.
func (pc *PaymentController) HandleWebhook(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		res.BadRequest(c, err)
		return
	}

	// Log the received transaction for debugging
	fmt.Printf("Received transaction: %+v\n", transaction)

	// Attempt to update the transaction status
	err := pc.service.UpdateTransactionStatus(&transaction)
	if err != nil {
		res.InternalServerError(c, err)
		return
	}

	// After the update, set the ID to the transaction ID
	transactionID := transaction.ID // Ensure this gets set in UpdateTransactionStatus
	res.SuccessOK(c, gin.H{
		"message":        "Transaction updated",
		"transaction_id": transactionID, // Return the correct ID
	})
}
