// middleware/jwt.go
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

var secretKey = []byte("your_secret_key") //----------

func AuthMiddleware() gin.HandlerFunc { //This function returns a middleware handler that checks the presence and validity of a JWT token in the Authorization header of incoming requests.
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ") //Extracts the token from the Authorization header, expecting it to be prefixed by "Bearer ".

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) { //Attempts to parse the token using the jwt.Parse() function with the predefined secretKey.
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort() //If the token is missing, invalid, or parsing fails, it returns a 401 Unauthorized status with an error message and stops further processing by calling c.Abort().
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("merchant_id", claims["id"]) //If the token is valid, it retrieves the claims (payload) from the token (in this case, merchant_id) and stores it in the request context (c.Set()).
		}

		c.Next()
	}
}
