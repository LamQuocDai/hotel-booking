package routes

import (
	"room-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupReviewRoutes(r *gin.Engine, h *handlers.ReviewHandler) {
	ReviewGroup := r.Group("/reviews")
	{
		ReviewGroup.GET("", h.GetAllReviews)
		ReviewGroup.GET("/:id", h.GetReviewByID)
		ReviewGroup.POST("", h.CreateReview)
		ReviewGroup.PUT("/:id", h.UpdatedReview)
		ReviewGroup.DELETE("/:id", h.DeleteReview)
	}
}
