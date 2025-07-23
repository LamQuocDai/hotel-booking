package routes

import (
	"my-app/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupAccountRoutes(r *gin.Engine, h *handlers.AccountHandler) {
	accountGroup := r.Group("/accounts")
	{
		accountGroup.GET("", h.GetAllAccounts)
		accountGroup.GET("/:id", h.GetAccountByID)
		accountGroup.POST("", h.CreateAccount)
		accountGroup.PUT("/:id", h.UpdatedAccount)
		accountGroup.DELETE("/:id", h.DeleteAccount)
	}
}
