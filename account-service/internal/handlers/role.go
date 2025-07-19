package handlers

import (
	"my-app/internal/models"
	"my-app/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RoleHandler struct {
	roleService *services.RoleService
	validate    *validator.Validate
}

func NewRoleHandler(roleService *services.RoleService) *RoleHandler {
	return &RoleHandler{roleService: roleService, validate: validator.New()}
}

func (h *RoleHandler) GetAllRoles(c *gin.Context) {
	roles, err := h.roleService.GetAllRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, roles)
}

func (h *RoleHandler) GetRoleByID(c *gin.Context) {
	id := c.Param("id")
	account, err := h.roleService.GetRoleByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, account)
}

func (h *RoleHandler) CreateRole(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.validate.Struct(role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.roleService.CreateRole(&role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, role)
}

func (h *RoleHandler) UpdateRole(c *gin.Context) {
	id := c.Param("id")
	var updatedRole models.Role
	if err := c.ShouldBindJSON(&updatedRole); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.validate.Struct(updatedRole); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.roleService.UpdateRole(id, &updatedRole); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedRole)
}

func (h *RoleHandler) DeleteRole(c *gin.Context) {
	id := c.Param("id")
	if err := h.roleService.DeleteRole(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
