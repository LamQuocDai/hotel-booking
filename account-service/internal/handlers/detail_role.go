package handlers

import (
	"my-app/internal/models"
	"my-app/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type DetailRoleHandler struct {
	detailRoleService *services.DetailRoleService
	validate          *validator.Validate
}

func NewDetailRoleHandler(detailRoleService *services.DetailRoleService) *DetailRoleHandler {
	return &DetailRoleHandler{detailRoleService: detailRoleService, validate: validator.New()}
}

func (h *DetailRoleHandler) GetAllDetailRoles(c *gin.Context) {
	detailRoles, err := h.detailRoleService.GetAllDetailRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Get all detail roles success!", "detailRoles": detailRoles})
}

func (h *DetailRoleHandler) GetDetailRoleByID(c *gin.Context) {
	id := c.Param("id")
	detailRole, err := h.detailRoleService.GetDetailRoleByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Get detail role by ID success!", "detailRoles": detailRole})
}

func (h *DetailRoleHandler) CreateDetailRole(c *gin.Context) {
	var detailRole models.DetailRole
	if err := c.ShouldBindJSON(&detailRole); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.validate.Struct(detailRole); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.detailRoleService.CreateDetailRole(&detailRole); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Created detail role success!", "detailRole": detailRole})
}

func (h *DetailRoleHandler) UpdatedDetailRole(c *gin.Context) {
	id := c.Param("id")
	var updatedDetailRole models.DetailRole
	if err := c.ShouldBindJSON(&updatedDetailRole); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.validate.Struct(updatedDetailRole); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.detailRoleService.UpdatedDetailRole(id, &updatedDetailRole); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Updated detail role success!", "updatedDetailRole": updatedDetailRole})
}

func (h *DetailRoleHandler) DeleteDetailRole(c *gin.Context) {
	id := c.Param("id")
	if err := h.detailRoleService.DeleteDetailRole(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted detail role success!"})
}
