package services

import (
	"decentralised_payment_gateway/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type PaymentService struct {
	db *gorm.DB
}

func NewPaymentService(db *gorm.DB) *PaymentService {
	return &PaymentService{db: db}
}

// CreatePayment creates a new payment request in the database.
func (ps *PaymentService) CreatePayment(payment *models.Payment) error {
	return ps.db.Create(payment).Error
}

// GetPaymentsForMerchant fetches payments for a specific merchant.
func (ps *PaymentService) GetPaymentsForMerchant(merchantID uint) ([]models.Payment, error) {
	var payments []models.Payment
	if err := ps.db.Where("merchant_id = ?", merchantID).Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}

// UpdateTransactionStatus updates a transaction's status based on its hash.
// UpdateTransactionStatus updates a transaction's status based on its hash.
func (ps *PaymentService) UpdateTransactionStatus(transaction *models.Transaction) error {
	fmt.Println("Attempting to update transaction with hash:", transaction.TransactionHash)

	// Find the existing transaction based on transaction_hash
	var existingTransaction models.Transaction
	err := ps.db.Where("transaction_hash = ?", transaction.TransactionHash).First(&existingTransaction).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Transaction not found for hash:", transaction.TransactionHash)
			// Create new transaction if not found
			transaction.ID = 0 // Ensure ID is 0 for new record
			if err := ps.db.Create(transaction).Error; err != nil {
				fmt.Println("Error creating transaction:", err)
				return err
			}
			fmt.Println("Transaction created successfully:", transaction)
			return nil
		}
		fmt.Println("Error finding transaction:", err)
		return err
	}

	// Log existing transaction before update
	fmt.Println("Existing transaction found before update:", existingTransaction)

	// Update the specific fields you want to modify
	existingTransaction.Amount = transaction.Amount         // Update amount
	existingTransaction.Status = transaction.Status         // Update status
	existingTransaction.Blockchain = transaction.Blockchain // Update blockchain if needed

	// Attempt to save the updated transaction
	result := ps.db.Save(&existingTransaction)

	if result.Error != nil {
		fmt.Println("Error during update:", result.Error)
		return result.Error
	}

	fmt.Println("Rows affected after update:", result.RowsAffected)

	if result.RowsAffected == 0 {
		fmt.Println("No rows updated. Check the update criteria.")
	} else {
		fmt.Println("Transaction updated successfully:", existingTransaction)
	}

	// Set the ID for the transaction after the update
	transaction.ID = existingTransaction.ID // Ensure to set the ID
	return nil
}
