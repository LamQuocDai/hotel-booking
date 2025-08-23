package routes

import (
	"room-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupServiceRoutes(router *gin.Engine, h *handlers.ServiceHandler) {
	serviceGroup := router.Group("/services")
	{
		serviceGroup.GET("/", h.GetAllServices)
		serviceGroup.GET("/:id", h.GetServiceByID)
		serviceGroup.POST("/", h.CreateService)
		serviceGroup.PUT("/:id", h.UpdateService)
		serviceGroup.DELETE("/:id", h.DeleteService)
	}

}
