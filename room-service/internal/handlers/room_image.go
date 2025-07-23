package handlers

import (
	"net/http"
	"room-service/internal/models"
	"room-service/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RoomImageHandler struct {
	roomImageService *services.RoomImageService
	validate         *validator.Validate
}

func NewRoomImageHandler(roomImageService *services.RoomImageService) *RoomImageHandler {
	return &RoomImageHandler{roomImageService: roomImageService, validate: validator.New()}
}

func (h *RoomImageHandler) GetAllRoomImages(c *gin.Context) {
	roomImages, err := h.roomImageService.GetAllRoomImages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, roomImages)
}

func (h *RoomImageHandler) GetRoomImageByID(c *gin.Context) {
	id := c.Param("id")
	roomImage, err := h.roomImageService.GetRoomImageByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, roomImage)
}

func (h *RoomImageHandler) CreateRoomImage(c *gin.Context) {
	var roomImage models.RoomImage
	if err := c.ShouldBindJSON(&roomImage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.validate.Struct(roomImage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.roomImageService.CreateRoomImage(&roomImage); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, roomImage)
}

func (h *RoomImageHandler) UpdatedRoomImage(c *gin.Context) {
	id := c.Param("id")
	var updatedRoomImage models.RoomImage
	if err := c.ShouldBindJSON(&updatedRoomImage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.validate.Struct(updatedRoomImage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.roomImageService.UpdatedRoomImage(id, &updatedRoomImage); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedRoomImage)
}

func (h *RoomImageHandler) DeleteRoomImage(c *gin.Context) {
	id := c.Param("id")
	if err := h.roomImageService.DeleteRoomImage(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
