package handlers

import (
	"net/http"
	"room-service/internal/models"
	"room-service/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ServiceHandler struct {
	serviceService *services.ServiceService
	validate       *validator.Validate
}

func NewServiceHandler(serviceService *services.ServiceService) *ServiceHandler {
	return &ServiceHandler{
		serviceService: serviceService,
		validate:       validator.New(),
	}
}

func (h *ServiceHandler) GetAllServices(c *gin.Context) {
	services, err := h.serviceService.GetAllServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Get all services successfully!", "services": services})
}

func (h *ServiceHandler) GetServiceByID(c *gin.Context) {
	id := c.Param("id")
	service, err := h.serviceService.GetServiceByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Get service by id successfully!", "service": service})
}

func (h *ServiceHandler) CreateService(c *gin.Context) {
	var service models.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.validate.Struct(service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.serviceService.CreateService(&service); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Create service successfully!", "service": service})
}

func (h *ServiceHandler) UpdateService(c *gin.Context) {
	id := c.Param("id")
	var updatedService models.Service
	if err := c.ShouldBindJSON(&updatedService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.validate.Struct(updatedService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.serviceService.UpdateService(id, &updatedService); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Update service successfully!", "service": updatedService})
}

func (h *ServiceHandler) DeleteService(c *gin.Context) {
	id := c.Param("id")
	if err := h.serviceService.DeleteService(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Delete service successfully!"})
}
