package handlers

import (
	"net/http"
	"service-service/internal/models"
	"service-service/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ServiceHandler struct {
	serviceService *services.ServiceService
	validate       *validator.Validate
}

func NewServiceHandler(serviceService *services.ServiceService) *ServiceHandler {
	return &ServiceHandler{serviceService: serviceService, validate: validator.New()}
}

func (h *ServiceHandler) GetAllServices(c *gin.Context) {
	services, err := h.serviceService.GetAllServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Get all services success!", "services": services})
}

func (h *ServiceHandler) GetServiceByID(c *gin.Context) {
	id := c.Param("id")
	service, err := h.serviceService.GetServiceByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Get service success!", "service": service})
}

func (h *ServiceHandler) CreateService(c *gin.Context) {
	var service models.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := h.serviceService.CreateService(&service); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Create service success!", "service": service})
}

func (h *ServiceHandler) UpdatedService(c *gin.Context) {
	id := c.Param("id")
	var updatedService models.Service
	if err := c.ShouldBindJSON(&updatedService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := h.serviceService.UpdatedService(id, &updatedService); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Updated service success!", "service": updatedService})
}

func (h *ServiceHandler) DeleleService(c *gin.Context) {
	id := c.Param("id")
	if err := h.serviceService.DeleteService(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted service success!"})
}
