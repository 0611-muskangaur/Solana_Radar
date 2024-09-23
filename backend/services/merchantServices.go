// services/merchantService.go
package services

import (
	"decentralised_payment_gateway/db"
	"decentralised_payment_gateway/models"
)

func CreateMerchant(merchant *models.Merchant) { // This function accepts a Merchant object and saves it to the database
	db.DB.Create(merchant)
}

func GetMerchantByWallet(walletAddress string) models.Merchant { //This function retrieves a Merchant record from the database based on a given wallet address.
	var merchant models.Merchant
	db.DB.Where("wallet_address = ?", walletAddress).First(&merchant)
	return merchant
}
