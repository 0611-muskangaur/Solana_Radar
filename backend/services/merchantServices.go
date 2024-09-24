package services

import (
	"decentralised_payment_gateway/db"
	"decentralised_payment_gateway/models"
	"fmt"
	"gorm.io/gorm"
)

// CreateMerchant saves a Merchant object to the database and returns an error if it fails
func CreateMerchant(merchant *models.Merchant) error {
	// Log the merchant data before saving
	fmt.Printf("Merchant before DB save:\n Name: %s, WalletAddress: %s, PreferredToken: %s\n",
		merchant.Name, merchant.WalletAddress, merchant.PreferredToken)

	err := db.DB.Create(merchant).Error
	if err != nil {
		// Check for unique constraint violation by examining the error message
		if err.Error() == "pq: duplicate key value violates unique constraint" {
			return fmt.Errorf("unique constraint violation for wallet address or password")
		}
		return fmt.Errorf("could not create merchant: %w", err)
	}
	return nil
}

// GetMerchantByWallet retrieves a Merchant record based on a wallet address
func GetMerchantByWallet(walletAddress string) (models.Merchant, error) {
	var merchant models.Merchant
	err := db.DB.Where("wallet_address = ?", walletAddress).First(&merchant).Error
	if err != nil {
		// Check if the error is a record not found error
		if err == gorm.ErrRecordNotFound {
			return merchant, fmt.Errorf("merchant not found: %w", err)
		}
		return merchant, fmt.Errorf("could not retrieve merchant: %w", err)
	}
	return merchant, nil
}
