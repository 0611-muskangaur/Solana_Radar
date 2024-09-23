// services/paymentService.go
package services

import (
	"decentralised_payment_gateway/db"
	"decentralised_payment_gateway/models"
)

func CreatePayment(payment *models.Payment) {
	db.DB.Create(payment)
}

func GetPaymentsForMerchant(merchantID uint) []models.Payment {
	var payments []models.Payment
	db.DB.Where("merchant_id = ?", merchantID).Find(&payments)
	return payments
}
