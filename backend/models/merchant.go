package models

import "time"

type Merchant struct {
	ID             uint      `gorm:"primaryKey"`
	Name           string    `gorm:"type:varchar(100);not null"`
	Email          string    `gorm:"type:varchar(100);unique;not null"`
	WalletAddress  string    `gorm:"type:varchar(100);not null"`
	PreferredToken string    `gorm:"type:varchar(50)"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
}
