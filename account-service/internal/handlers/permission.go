package handlers

import (
	"my-app/internal/models"
	"my-app/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PermissionHandler struct {
	permissionService *services.PermissionService
	validate          *validator.Validate
}

func NewPermissionHandler(permissionService *services.PermissionService) *PermissionHandler {
	return &PermissionHandler{permissionService: permissionService, validate: validator.New()}
}

func (h *PermissionHandler) GetAllPermissions(c *gin.Context) {
	permissions, err := h.permissionService.GetAllPermissions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Get all permission success!", "permissions": permissions})
}

func (h *PermissionHandler) GetPermissionByID(c *gin.Context) {
	id := c.Param("id")
	permission, err := h.permissionService.GetPermissionByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Get permission by ID success!", "permission": permission})
}

func (h *PermissionHandler) CreatePermission(c *gin.Context) {
	var permission models.Permission
	if err := c.ShouldBindJSON(&permission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.validate.Struct(permission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.permissionService.CreatePermission(&permission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Create permission success!", "permission": permission})
}

func (h *PermissionHandler) UpdatedPermission(c *gin.Context) {
	id := c.Param("id")
	var updatedPermission models.Permission
	if err := c.ShouldBindJSON(&updatedPermission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.validate.Struct(updatedPermission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.permissionService.UpdatedPermission(id, &updatedPermission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Udpated permission success!", "updatedPermission": updatedPermission})
}

func (h *PermissionHandler) DeletePermission(c *gin.Context) {
	id := c.Param("id")
	if err := h.permissionService.DeletePermission(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted permission success!"})
}
