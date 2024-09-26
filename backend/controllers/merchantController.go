package controllers

import (
	"net/http"
	"time"

	"decentralised_payment_gateway/models"
	"decentralised_payment_gateway/services"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("your_secret_key")

func RegisterMerchant(c *gin.Context) {
	var input models.Merchant
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure that wallet address and preferred token are provided
	if input.WalletAddress == "" || input.PreferredToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wallet address and preferred token are required"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}
	input.Password = string(hashedPassword)

	// Create the merchant and handle any errors
	if err := services.CreateMerchant(&input); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Merchant registered successfully"})
}

// Merchant login
func LoginMerchant(c *gin.Context) {
	var input models.MerchantLogin
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Retrieve the merchant by wallet address
	merchant, err := services.GetMerchantByWallet(input.WalletAddress)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Compare the provided password with the hashed password stored in the database
	if bcrypt.CompareHashAndPassword([]byte(merchant.Password), []byte(input.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  merchant.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(), // Token expiration (72 hours)
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Respond with JWT token
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
