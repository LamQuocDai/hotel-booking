package routes

import (
	"room-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoomBookingRoutes(r *gin.Engine, h *handlers.RoomBookingHandler) {
	roomBookingGroup := r.Group("/room-bookings")
	{
		roomBookingGroup.GET("", h.GetAllRoomBookings)
		roomBookingGroup.GET("/:id", h.GetRoomBookingByID)
		roomBookingGroup.POST("", h.CreateRoomBooking)
		roomBookingGroup.PUT("/:id", h.UpdatedRoomBooking)
		roomBookingGroup.DELETE("/:id", h.DeleteRoomBooking)
	}
}
