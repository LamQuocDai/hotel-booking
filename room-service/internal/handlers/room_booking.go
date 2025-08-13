package handlers

import (
	"net/http"
	"room-service/internal/models"
	"room-service/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RoomBookingHandler struct {
	roomBookingService *services.RoomBookingService
	validate           *validator.Validate
}

func NewRoomBookingHandler(roomBookingService *services.RoomBookingService) *RoomBookingHandler {
	return &RoomBookingHandler{roomBookingService: roomBookingService, validate: validator.New()}
}

func (h *RoomBookingHandler) GetAllRoomBookings(c *gin.Context) {
	roomBookings, err := h.roomBookingService.GetAllRoomBookings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Get all room bookings success!", "roomBookings": roomBookings})
}

func (h *RoomBookingHandler) GetRoomBookingByID(c *gin.Context) {
	id := c.Param("id")
	roomBooking, err := h.roomBookingService.GetRoomBookingByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Get room booking by id success!", "roomBooking": roomBooking})
}

func (h *RoomBookingHandler) CreateRoomBooking(c *gin.Context) {
	var roomBooking models.RoomBooking
	if err := c.ShouldBindJSON(&roomBooking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := h.validate.Struct(roomBooking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := h.roomBookingService.CreateRoomBooking(&roomBooking); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Created room booking success!", "roomBooking": roomBooking})
}

func (h *RoomBookingHandler) UpdatedRoomBooking(c *gin.Context) {
	id := c.Param("id")
	var updatedRoomBooking models.RoomBooking
	if err := c.ShouldBindJSON(&updatedRoomBooking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := h.validate.Struct(updatedRoomBooking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := h.roomBookingService.UpdatedRoomBooking(id, &updatedRoomBooking); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Updated room booking success!", "updatedRoomBooking": updatedRoomBooking})
}

func (h *RoomBookingHandler) DeleteRoomBooking(c *gin.Context) {
	id := c.Param("id")
	if err := h.roomBookingService.DeleteRoomBooking(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted room booking success!"})
}
