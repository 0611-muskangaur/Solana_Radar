package webhooks

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func WebhookHandler(c *gin.Context) {
	var payload map[string]interface{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Process the webhook payload (e.g., update payment status)
	// Assume you have a service to handle this logic

	c.JSON(http.StatusOK, gin.H{"message": "Webhook received"})
}
