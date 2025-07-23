package routes

import (
	"room-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupLocationRoute(r *gin.Engine, h *handlers.LocationHandler) {
	locationGroup := r.Group("/locations")
	{
		locationGroup.GET("", h.GetAllLocations)
		locationGroup.GET("/:id", h.GetLocationByID)
		locationGroup.POST("", h.CreateLocation)
		locationGroup.PUT("/:id", h.UpdatedLocation)
		locationGroup.DELETE("/:id", h.DeleteLocation)
	}
}
