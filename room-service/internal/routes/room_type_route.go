package routes

import (
	"room-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoomTypeRouter(r *gin.Engine, h *handlers.RoomTypeHandler) {
	RoomTypeGroup := r.Group("/room-types")
	{
		RoomTypeGroup.GET("", h.GetAllRoomTypes)
		RoomTypeGroup.GET("/:id", h.GetRoomTypeByID)
		RoomTypeGroup.POST("", h.CreateRoomType)
		RoomTypeGroup.PUT("/:id", h.UpdatedRoomType)
		RoomTypeGroup.DELETE("/:id", h.DeleteRoomType)
	}
}
