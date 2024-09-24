package models

import "time"

// Merchant represents a merchant in the system
type Merchant struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	WalletAddress  string    `json:"wallet_address"` // ensure the json tag is here
	Password       string    `json:"password"`
	PreferredToken string    `json:"preferred_token"` // ensure the json tag is here
	CreatedAt      time.Time `json:"created_at"`
}

// MerchantLogin represents the login credentials for a merchant
type MerchantLogin struct {
	WalletAddress string `json:"wallet_address"`
	Password      string `json:"password"`
}
