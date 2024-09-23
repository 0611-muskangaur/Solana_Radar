// controllers/merchantController.go
package controllers

import (
	"net/http"
	"time"

	"decentralised_payment_gateway/models"
	"decentralised_payment_gateway/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("your_secret_key")

// Merchant registration
func RegisterMerchant(c *gin.Context) {
	var input models.Merchant
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	input.Password = string(hashedPassword)

	services.CreateMerchant(&input)
	c.JSON(http.StatusCreated, gin.H{"message": "Merchant registered successfully"})
}

// Merchant login
func LoginMerchant(c *gin.Context) {
	var input models.MerchantLogin
	var merchant models.Merchant

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	merchant = services.GetMerchantByWallet(input.WalletAddress)
	if merchant.ID == 0 || bcrypt.CompareHashAndPassword([]byte(merchant.Password), []byte(input.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  merchant.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
