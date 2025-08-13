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
	roomBookingService := services.NewRoomBookingService(db)
	roomBookingHandler := handlers.NewRoomBookingHandler(roomBookingService)

	//Register
	SetupLocationRoutes(r, locationHandler)
	SetupRoomTypeRoutes(r, roomTypeHandler)
	SetupRoomRoutes(r, roomHandler)
	SetupRoomImageRoutes(r, roomImageHandler)
	SetupReviewRoutes(r, reviewHandler)
	SetupRoomBookingRoutes(r, roomBookingHandler)

	return r
}
