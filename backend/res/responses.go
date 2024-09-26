package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SuccessResponse is a general structure for API success responses.
func SuccessResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{
		"status":  statusCode,
		"message": "success",
		"data":    data,
	})
}

// SuccessCreated is a 201 Created response.
func SuccessCreated(c *gin.Context, data interface{}) {
	SuccessResponse(c, http.StatusCreated, data)
}

// SuccessOK is a 200 OK response.
func SuccessOK(c *gin.Context, data interface{}) {
	SuccessResponse(c, http.StatusOK, data)
}
