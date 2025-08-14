package routes

import (
	"my-app/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupPermissionRoutes(r *gin.Engine, h *handlers.PermissionHandler) {
	permissionGroup := r.Group("/permissions")
	{
		permissionGroup.GET("", h.GetAllPermissions)
		permissionGroup.GET("/:id", h.GetPermissionByID)
		permissionGroup.POST("", h.CreatePermission)
		permissionGroup.PUT("/:id", h.UpdatedPermission)
		permissionGroup.DELETE("/:id", h.DeletePermission)
	}
}
