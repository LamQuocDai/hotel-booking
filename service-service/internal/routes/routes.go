package routes

import (
	"service-service/internal/handlers"
	"service-service/internal/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(db *mongo.Database) *gin.Engine {
	r := gin.Default()
	//
	serviceService := services.NewServiceService(db)
	serviceHandler := handlers.NewServiceHandler(serviceService)

	//
	SetupServiceRoutes(r, serviceHandler)
	return r
}
