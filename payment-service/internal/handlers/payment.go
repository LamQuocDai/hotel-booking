package handlers

import (
	"net/http"
	"payment-service/internal/models"
	"payment-service/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PaymentHandler struct {
	paymentService *services.PaymentService
	validate       *validator.Validate
}

func NewPaymentHandler(paymentService *services.PaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService: paymentService, validate: validator.New()}
}

func (h *PaymentHandler) GetAllPayments(c *gin.Context) {
	payments, err := h.paymentService.GetAllPayments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Get all payments success!", "payments": payments})
}

func (h *PaymentHandler) GetPaymentByID(c *gin.Context) {
	id := c.Param("id")
	payment, err := h.paymentService.GetPaymentByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Get payment by ID success!", "payment": payment})
}

func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var payment models.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := h.paymentService.CreatePayment(&payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Created payment success!", "payment": payment})
}

func (h *PaymentHandler) UpdatedPayment(c *gin.Context) {
	id := c.Param("id")
	var updatedPayment models.Payment
	if err := c.ShouldBindJSON(&updatedPayment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := h.paymentService.UpdatedPayment(id, &updatedPayment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Updated payment success!", "payment": updatedPayment})
}

func (h *PaymentHandler) DeletePayment(c *gin.Context) {
	id := c.Param("id")
	if err := h.paymentService.DeletePayment(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted payment success!"})
}
