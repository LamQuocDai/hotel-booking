package routes

import (
	"my-app/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupDetailRoleRoutes(r *gin.Engine, h *handlers.DetailRoleHandler) {
	detailRoleGroup := r.Group("/detailRoles")
	{
		detailRoleGroup.GET("", h.GetAllDetailRoles)
		detailRoleGroup.GET("/:id", h.GetDetailRoleByID)
		detailRoleGroup.POST("", h.CreateDetailRole)
		detailRoleGroup.PUT("/:id", h.UpdatedDetailRole)
		detailRoleGroup.DELETE("/:id", h.DeleteDetailRole)
	}
}
