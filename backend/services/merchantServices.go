// services/merchantService.go
package services

import (
	"decentralised_payment_gateway/db"
	"decentralised_payment_gateway/models"
)

func CreateMerchant(merchant *models.Merchant) {
	db.DB.Create(merchant)
}

func GetMerchantByWallet(walletAddress string) models.Merchant {
	var merchant models.Merchant
	db.DB.Where("wallet_address = ?", walletAddress).First(&merchant)
	return merchant
}
