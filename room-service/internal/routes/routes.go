package routes

import (
	"room-service/internal/handlers"
	"room-service/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Initialize
	locationService := services.NewLocationService(db)
	locationHandler := handlers.NewLocationHandler(locationService)
	roomTypeService := services.NewRoomTypeService(db)
	roomTypeHandler := handlers.NewRoomTypeHandler(roomTypeService)
	roomService := services.NewRoomService(db)
	roomHandler := handlers.NewRoomHandler(roomService)
	roomImageService := services.NewRoomImageService(db)
	roomImageHandler := handlers.NewRoomImageHandler(roomImageService)
	reviewService := services.NewReviewService(db)
	reviewHandler := handlers.NewReviewHandler(reviewService)

	//Register
	SetupLocationRoute(r, locationHandler)
	SetupRoomTypeRouter(r, roomTypeHandler)
	SetupRoomRouter(r, roomHandler)
	SetupRoomImageRouter(r, roomImageHandler)
	SetupReviewRouter(r, reviewHandler)

	return r
}
