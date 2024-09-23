// main.go
package main

import (
	"decentralised_payment_gateway/config"
	"decentralised_payment_gateway/db"
	"decentralised_payment_gateway/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	db.ConnectPostgres()

	r := gin.Default()

	routes.MerchantRoutes(r)
	routes.PaymentRoutes(r)

	r.Run(":8080")
}
