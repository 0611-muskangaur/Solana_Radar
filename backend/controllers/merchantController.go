// controllers/merchantController.go
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

var secretKey = []byte("your_secret_key") //------------------

// Merchant registration
func RegisterMerchant(c *gin.Context) {
	var input models.Merchant //Receives a JSON payload representing a new merchant (models.Merchant)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	input.Password = string(hashedPassword) //It hashes the merchant's password using bcrypt to ensure passwords are securely stored.

	services.CreateMerchant(&input) //The hashed password is then stored, and the merchant is created via the services.CreateMerchant function.
	c.JSON(http.StatusCreated, gin.H{"message": "Merchant registered successfully"}) //Returns an HTTP 201 Created status and a success message if registration is successful, or a 400 Bad Request status if there's an issue with the input.
}

// Merchant login
func LoginMerchant(c *gin.Context) {
	var input models.MerchantLogin //Expects a JSON payload containing login credentials (models.MerchantLogin), including the wallet address and password.
	var merchant models.Merchant

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	merchant = services.GetMerchantByWallet(input.WalletAddress)  //Retrieves the merchant from the database by wallet address
	if merchant.ID == 0 || bcrypt.CompareHashAndPassword([]byte(merchant.Password), []byte(input.Password)) != nil { //Compares the provided password with the hashed password stored in the database
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  merchant.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(), //If the credentials are valid, a JWT is generated containing the merchant's ID and an expiration time (72 hours from login).
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})   //Returns the JWT as a token in the response if the login is successful, or an error message with 401 Unauthorized for invalid credentials.
		return     
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
