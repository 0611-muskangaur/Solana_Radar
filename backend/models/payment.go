package models

import "time"

type Payment struct {
	ID              uint      `gorm:"primaryKey"`
	MerchantID      uint      `gorm:"not null"`
	Amount          float64   `gorm:"type:decimal(18,8);not null"`
	Currency        string    `gorm:"type:varchar(10)"`
	PaymentStatus   string    `gorm:"type:varchar(50);default:'pending'"`
	TransactionHash string    `gorm:"type:varchar(255);unique"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
}

type Transaction struct {
	ID               uint      `gorm:"primaryKey"`
	PaymentRequestID uint      `gorm:"not null"`
	Amount           float64   `gorm:"type:decimal(18,8);not null"`
	Status           string    `gorm:"type:varchar(50);default:'pending'"`
	Blockchain       string    `gorm:"type:varchar(50);not null"`
	Date             time.Time `gorm:"autoCreateTime"`
	TransactionHash  string    `gorm:"type:varchar(255);unique"`
}
