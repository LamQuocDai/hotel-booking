package routes

import (
	"payment-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupPromotionRoutes(r *gin.Engine, h *handlers.PromotionHandler) {
	promotionGroup := r.Group("/promotions")
	{
		promotionGroup.GET("", h.GetAllPromotions)
		promotionGroup.GET("/:id", h.GetPromotionByID)
		promotionGroup.POST("", h.CreatePromotion)
		promotionGroup.PUT("/:id", h.UpdatedPromotion)
		promotionGroup.DELETE("/:id", h.DeletePromotion)
	}
}
