package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HandleError formats error responses for the API.
func HandleError(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, gin.H{
		"status":  statusCode,
		"message": err.Error(),
	})
	c.Abort()
}

// BadRequest handles 400 Bad Request errors.
func BadRequest(c *gin.Context, err error) {
	HandleError(c, http.StatusBadRequest, err)
}

// Unauthorized handles 401 Unauthorized errors.
func Unauthorized(c *gin.Context, err error) {
	HandleError(c, http.StatusUnauthorized, err)
}

// InternalServerError handles 500 Internal Server errors.
func InternalServerError(c *gin.Context, err error) {
	HandleError(c, http.StatusInternalServerError, err)
}
