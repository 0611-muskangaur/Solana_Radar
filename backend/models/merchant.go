// models/merchant.go
package models

type Merchant struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"size:100"`
	WalletAddress  string `gorm:"size:100;not null"`
	PreferredToken string `gorm:"size:50"`
	Password       string `gorm:"size:255;not null"`
}

type MerchantLogin struct {
	WalletAddress string `json:"wallet_address"`
	Password      string `json:"password"`
}
