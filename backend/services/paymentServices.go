package services

import (
	"decentralised_payment_gateway/models"
	"gorm.io/gorm"
)

type PaymentService struct {
	db *gorm.DB
}

func NewPaymentService(db *gorm.DB) *PaymentService {
	return &PaymentService{db: db}
}

// Create a payment.
func (ps *PaymentService) CreatePayment(payment *models.Payment) error {
	return ps.db.Create(payment).Error
}

// Fetch payments for a merchant.
func (s *PaymentService) GetPaymentsForMerchant(merchantID uint) ([]models.Payment, error) {
	var payments []models.Payment
	if err := s.db.Where("merchant_id = ?", merchantID).Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}

// Create a transaction.
func (ps *PaymentService) CreateTransaction(transaction *models.Transaction) error {
	return ps.db.Create(transaction).Error
}

// Update transaction status based on webhook data.
func (ps *PaymentService) UpdateTransactionStatus(transaction *models.Transaction) error {
	return ps.db.Model(&models.Transaction{}).
		Where("transaction_hash = ?", transaction.TransactionHash).
		Update("status", transaction.Status).Error
}
