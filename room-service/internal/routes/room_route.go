package routes

import (
	"room-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoomRoutes(r *gin.Engine, h *handlers.RoomHandler) {
	RoomGroup := r.Group("/rooms")
	{
		RoomGroup.GET("", h.GetAllRooms)
		RoomGroup.GET("/:id", h.GetRoomByID)
		RoomGroup.POST("", h.CreateRoom)
		RoomGroup.PUT("/:id", h.UpdatedRoom)
		RoomGroup.DELETE("/:id", h.DeleteRoom)
	}
}
