package handlers

import (
	"net/http"
	"room-service/internal/models"
	"room-service/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RoomTypeHandler struct {
	roomTypeService *services.RoomTypeService
	validate        *validator.Validate
}

func NewRoomTypeHandler(roomTypeService *services.RoomTypeService) *RoomTypeHandler {
	return &RoomTypeHandler{roomTypeService: roomTypeService, validate: validator.New()}
}

func (h *RoomTypeHandler) GetAllRoomTypes(c *gin.Context) {
	roomTypes, err := h.roomTypeService.GetAllRoomTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, roomTypes)
}

func (h *RoomTypeHandler) GetRoomTypeByID(c *gin.Context) {
	id := c.Param("id")
	roomType, err := h.roomTypeService.GetRoomTypeByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, roomType)
}

func (h *RoomTypeHandler) CreateRoomType(c *gin.Context) {
	var roomType models.RoomType
	if err := c.ShouldBindJSON(&roomType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.validate.Struct(roomType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.roomTypeService.CreateRoomType(&roomType); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, roomType)
}

func (h *RoomTypeHandler) UpdatedRoomType(c *gin.Context) {
	id := c.Param("id")
	var updatedRoomType models.RoomType
	if err := c.ShouldBindJSON(&updatedRoomType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.validate.Struct(updatedRoomType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.roomTypeService.UpdatedRoomType(id, &updatedRoomType); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedRoomType)
}

func (h *RoomTypeHandler) DeleteRoomType(c *gin.Context) {
	id := c.Param("id")
	if err := h.roomTypeService.DeleteRoomType(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
