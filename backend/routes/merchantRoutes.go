// routes/merchantRoutes.go
package routes

import (
	"decentralised_payment_gateway/controllers"

	"github.com/gin-gonic/gin"
)

func MerchantRoutes(r *gin.Engine) {
	r.POST("/register", controllers.RegisterMerchant)
	r.POST("/login", controllers.LoginMerchant)
}
