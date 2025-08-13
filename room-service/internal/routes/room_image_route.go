package routes

import (
	"room-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoomImageRoutes(r *gin.Engine, h *handlers.RoomImageHandler) {
	RoomImageGroup := r.Group("/room-images")
	{
		RoomImageGroup.GET("", h.GetAllRoomImages)
		RoomImageGroup.GET("/:id", h.GetRoomImageByID)
		RoomImageGroup.POST("", h.CreateRoomImage)
		RoomImageGroup.PUT("/:id", h.UpdatedRoomImage)
		RoomImageGroup.DELETE("/:id", h.DeleteRoomImage)
	}
}
