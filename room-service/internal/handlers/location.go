package handlers

import (
	"net/http"
	"room-service/internal/models"
	"room-service/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type LocationHandler struct {
	locationService *services.LocationService
	validate        *validator.Validate
}

func NewLocationHandler(locationService *services.LocationService) *LocationHandler {
	return &LocationHandler{locationService: locationService, validate: validator.New()}
}

func (h *LocationHandler) GetAllLocations(c *gin.Context) {
	locations, err := h.locationService.GetAllLocations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, locations)
}

func (h *LocationHandler) GetLocationByID(c *gin.Context) {
	id := c.Param("id")
	location, err := h.locationService.GetLocationByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, location)
}

func (h *LocationHandler) CreateLocation(c *gin.Context) {
	var location models.Location
	if err := c.ShouldBindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.validate.Struct(location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.locationService.CreateLocation(&location); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, location)
}

func (h *LocationHandler) UpdatedLocation(c *gin.Context) {
	id := c.Param("id")
	var updatedLocation models.Location
	if err := c.ShouldBindJSON(&updatedLocation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.validate.Struct(updatedLocation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.locationService.UpdatedLocation(id, &updatedLocation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedLocation)
}

func (h *LocationHandler) DeleteLocation(c *gin.Context) {
	id := c.Param("id")
	if err := h.locationService.DeleteLocation(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
