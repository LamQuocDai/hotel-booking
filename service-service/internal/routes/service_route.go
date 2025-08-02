package routes

import (
	"service-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupServiceRoutes(r *gin.Engine, h *handlers.ServiceHandler) {
	serviceGroup := r.Group("/services")
	{
		serviceGroup.GET("", h.GetAllServices)
		serviceGroup.GET("/:id", h.GetServiceByID)
		serviceGroup.POST("", h.CreateService)
		serviceGroup.PUT("/:id", h.UpdatedService)
		serviceGroup.DELETE("/:id", h.DeleleService)
	}
}
