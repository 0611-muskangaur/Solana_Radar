// services/paymentService.go
package services

import (
	"decentralised_payment_gateway/db"
	"decentralised_payment_gateway/models"
)

func CreatePayment(payment *models.Payment) { //This function accepts a Payment object and stores it in the database
	db.DB.Create(payment)
}

func GetPaymentsForMerchant(merchantID uint) []models.Payment { //This function retrieves all payments associated with a specific merchant, identified by their merchantID
	var payments []models.Payment
	db.DB.Where("merchant_id = ?", merchantID).Find(&payments)
	return payments
}
