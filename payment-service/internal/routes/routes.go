package routes

import (
	"payment-service/internal/handlers"
	"payment-service/internal/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(db *mongo.Database) *gin.Engine {
	r := gin.Default()

	//
	paymentService := services.NewPaymentService(db)
	paymentHandler := handlers.NewPaymentHandler(paymentService)
	promotionService := services.NewPromotionService(db)
	promotionHandler := handlers.NewPromotionHandler(promotionService)

	//
	SetupPaymentRoute(r, paymentHandler)
	SetupPromotionRoute(r, promotionHandler)

	return r
}
