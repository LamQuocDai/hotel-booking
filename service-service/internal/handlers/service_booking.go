package handlers

import (
	"net/http"
	"service-service/internal/models"
	"service-service/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ServiceBookingHandler struct {
	serviceBookingService *services.ServiceBookingService
	validate              *validator.Validate
}

func NewServiceBookingHandler(serviceBookingService *services.ServiceBookingService) *ServiceBookingHandler {
	return &ServiceBookingHandler{serviceBookingService: serviceBookingService, validate: validator.New()}
}

func (h *ServiceBookingHandler) GetAllServiceBookings(c *gin.Context) {
	serviceBookings, err := h.serviceBookingService.GetAllServiceBookings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Get all service booking success!", "serviceBookings": serviceBookings})
}

func (h *ServiceBookingHandler) GetServiceBookingByID(c *gin.Context) {
	id := c.Param("id")
	serviceBooking, err := h.serviceBookingService.GetServiceBookingByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Get service booking by id", "serviceBooking": serviceBooking})
}

func (h *ServiceBookingHandler) CreateServiceBooking(c *gin.Context) {
	var serviceBooking models.ServiceBooking
	if err := c.ShouldBindJSON(&serviceBooking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := h.serviceBookingService.CreateServiceBooking(&serviceBooking); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Created service booking success!", "serviceBooking": serviceBooking})
}

func (h *ServiceBookingHandler) UpdatedServiceBooking(c *gin.Context) {
	id := c.Param("id")
	var updatedServiceBooking models.ServiceBooking
	if err := c.ShouldBindJSON(&updatedServiceBooking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := h.serviceBookingService.UpdatedServiceBooking(id, &updatedServiceBooking); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Updated service booking success!", "updatedServiceBooking": updatedServiceBooking})
}

func (h *ServiceBookingHandler) DeleteServiceBooking(c *gin.Context) {
	id := c.Param("id")
	if err := h.serviceBookingService.DeleteServiceBooking(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted service booking success!"})
}
