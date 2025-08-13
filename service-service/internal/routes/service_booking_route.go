package routes

import (
	"service-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupServiceBookingRoutes(r *gin.Engine, h *handlers.ServiceBookingHandler) {
	serviceBookingGroup := r.Group("/service-bookings")
	{
		serviceBookingGroup.GET("", h.GetAllServiceBookings)
		serviceBookingGroup.GET("/:id", h.GetServiceBookingByID)
		serviceBookingGroup.POST("", h.CreateServiceBooking)
		serviceBookingGroup.PUT("/:id", h.UpdatedServiceBooking)
		serviceBookingGroup.DELETE("/:id", h.DeleteServiceBooking)
	}
}
