package routes

import (
	"my-app/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoleRoutes(r *gin.Engine, h *handlers.RoleHandler) {
	roleGroup := r.Group("/roles")
	{
		roleGroup.GET("", h.GetAllRoles)
		roleGroup.GET("/:id", h.GetRoleByID)
		roleGroup.POST("", h.CreateRole)
		roleGroup.PUT("/:id", h.UpdateRole)
		roleGroup.DELETE("/:id", h.DeleteRole)
	}
}
