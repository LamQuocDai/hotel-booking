package routes

import (
	"my-app/internal/handlers"
	"my-app/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Initialize services and handlers
	accountService := services.NewAccountService(db)
	accountHandler := handlers.NewAccountHandler(accountService)
	roleService := services.NewRoleService(db)
	roleHandler := handlers.NewRoleHandler(roleService)

	// Register route groups
	SetupAccountRoutes(r, accountHandler)
	SetupRoleRoutes(r, roleHandler)

	return r
}
