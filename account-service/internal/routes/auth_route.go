package routes

import (
	"my-app/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.Engine, h *handlers.AuthHandler) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", h.Login)
	}
}
