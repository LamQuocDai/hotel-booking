package handlers

import (
	"net/http"
	"payment-service/internal/models"
	"payment-service/internal/services"
	"payment-service/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PromotionHandler struct {
	promotionService *services.PromotionService
	validate         *validator.Validate
}

type CreatePromotionReq struct {
	Code        string `json:"code" binding:"required"`
	Description string `json:"description"`
	Discount    int    `json:"discount" binding:"required,min=0,max=100"`
	StartDate   string `json:"startDate" binding:"required"` // e.g. 2025-07-07 or 7/7/2025
	EndDate     string `json:"endDate" binding:"required"`
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
	var req CreatePromotionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	startT, err := utils.ParseFlexibleDate(req.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid startDate format"})
		return
	}
	endT, err := utils.ParseFlexibleDate(req.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid endDate format"})
		return
	}
	promotion := models.Promotion{
		Code:        req.Code,
		Description: req.Description,
		Discount:    req.Discount,
		StartDay:    utils.DateToDayInt(startT),
		EndDay:      utils.DateToDayInt(endT),
	}
	if err := h.promotionService.CreatePromotion(&promotion); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Created promotion success!", "promotion": promotion})
}

func (h *PromotionHandler) UpdatedPromotion(c *gin.Context) {
	id := c.Param("id")
	var req CreatePromotionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	startT, err := utils.ParseFlexibleDate(req.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid startDate format"})
		return
	}
	endT, err := utils.ParseFlexibleDate(req.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid endDate format"})
		return
	}
	updatedPromotion := models.Promotion{
		Code:        req.Code,
		Description: req.Description,
		Discount:    req.Discount,
		StartDay:    utils.DateToDayInt(startT),
		EndDay:      utils.DateToDayInt(endT),
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
