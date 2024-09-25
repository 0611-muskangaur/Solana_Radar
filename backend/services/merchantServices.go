package services

import (
	"decentralised_payment_gateway/db"
	"decentralised_payment_gateway/models"
	"fmt"
	"gorm.io/gorm"
)

func CreateMerchant(merchant *models.Merchant) error {
	// Check if the wallet address already exists
	var existingWalletMerchant models.Merchant
	err := db.DB.Where("wallet_address = ?", merchant.WalletAddress).First(&existingWalletMerchant).Error
	if err == nil {
		return fmt.Errorf("wallet address already exists")
	} else if err != gorm.ErrRecordNotFound {
		return fmt.Errorf("could not check wallet address: %w", err)
	}

	// Proceed to save the merchant
	err = db.DB.Create(merchant).Error
	if err != nil {
		return fmt.Errorf("could not create merchant: %w", err)
	}
	return nil
}

func GetMerchantByWallet(walletAddress string) (models.Merchant, error) {
	var merchant models.Merchant
	err := db.DB.Where("wallet_address = ?", walletAddress).First(&merchant).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return merchant, fmt.Errorf("merchant not found: %w", err)
		}
		return merchant, fmt.Errorf("could not retrieve merchant: %w", err)
	}
	return merchant, nil
}
