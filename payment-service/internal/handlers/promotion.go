package handlers

import (
	"net/http"
	"payment-service/internal/models"
	"payment-service/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PromotionHandler struct {
	promotionService *services.PromotionService
	validate         *validator.Validate
}

func NewPromotionHandler(promotionService *services.PromotionService) *PromotionHandler {
	return &PromotionHandler{promotionService: promotionService, validate: validator.New()}
}

func (h *PromotionHandler) GetAllPromotions(c *gin.Context) {
	promotions, err := h.promotionService.GetAllPromotions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Get all promotions success!", "promotions": promotions})
}

func (h *PromotionHandler) GetPromotionByID(c *gin.Context) {
	id := c.Param("id")
	promotion, err := h.promotionService.GetPromotionByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Get promotion by ID success!", "promotion": promotion})
}

func (h *PromotionHandler) CreatePromotion(c *gin.Context) {
	var promotion models.Promotion
	if err := c.ShouldBindJSON(&promotion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := h.promotionService.CreatePromotion(&promotion); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Created promotion success!", "promotion": promotion})
}

func (h *PromotionHandler) UpdatedPromotion(c *gin.Context) {
	id := c.Param("id")
	var updatedPromotion models.Promotion
	if err := c.ShouldBindJSON(&updatedPromotion); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if err := h.promotionService.UpdatedPromotion(id, &updatedPromotion); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Updated promotion success!", "promotion": updatedPromotion})
}

func (h *PromotionHandler) DeletePromotion(c *gin.Context) {
	id := c.Param("id")
	if err := h.promotionService.DeletePromotion(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted promotion success!"})
}
