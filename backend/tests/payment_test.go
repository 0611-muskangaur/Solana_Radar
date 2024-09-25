package tests

import (
	"testing"

	"decentralised_payment_gateway/models"
	"decentralised_payment_gateway/services"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupDB() *gorm.DB {
	dsn := "user=postgres password=Muskan0611@$ dbname=payment_gateway host=127.0.0.1 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&models.Payment{}, &models.Transaction{}) // Migrate both tables
	return db
}

func TestCreatePayment(t *testing.T) {
	db := setupDB()
	service := services.NewPaymentService(db)

	payment := &models.Payment{
		MerchantID:      1,
		Amount:          100.50,
		Currency:        "USD",
		PaymentStatus:   "pending",
		TransactionHash: "hash123",
	}

	err := service.CreatePayment(payment)
	assert.Nil(t, err)
	assert.NotZero(t, payment.ID) // Check if ID is set
}

func TestGetPaymentsForMerchant(t *testing.T) {
	db := setupDB()
	service := services.NewPaymentService(db)

	merchantID := uint(1)
	payments, err := service.GetPaymentsForMerchant(merchantID)
	assert.Nil(t, err)
	assert.IsType(t, []models.Payment{}, payments)
}

func TestWebhookUpdateTransactionStatus(t *testing.T) {
	db := setupDB()
	service := services.NewPaymentService(db)

	transaction := &models.Transaction{
		PaymentRequestID: 1,
		Amount:           200.50,
		Status:           "completed",
		Blockchain:       "Solana",
		TransactionHash:  "hash123",
	}

	err := service.UpdateTransactionStatus(transaction)
	assert.Nil(t, err)
}
