package services

import (
	"room-service/internal/models"

	"gorm.io/gorm"
)

type RoomBookingService struct {
	db *gorm.DB
}

func NewRoomBookingService(db *gorm.DB) *RoomBookingService {
	return &RoomBookingService{db: db}
}

func (s *RoomBookingService) GetAllRoomBookings() ([]models.RoomBooking, error) {
	var roomBookings []models.RoomBooking
	return roomBookings, s.db.Preload("Room").Find(&roomBookings).Error
}

func (s *RoomBookingService) GetRoomBookingByID(id string) (*models.RoomBooking, error) {
	var roomBooking models.RoomBooking
	return &roomBooking, s.db.Preload("Room").Where("id = ?", id).First(&roomBooking).Error
}

func (s *RoomBookingService) CreateRoomBooking(roomBooking *models.RoomBooking) error {
	return s.db.Create(*roomBooking).Error
}

func (s *RoomBookingService) UpdatedRoomBooking(id string, updatedRoomBooking *models.RoomBooking) error {
	var roomBooking models.RoomBooking
	if err := s.db.Where("id = ?", id).First(&roomBooking).Error; err != nil {
		return err
	}
	roomBooking.BookingId = updatedRoomBooking.BookingId
	roomBooking.RoomId = updatedRoomBooking.RoomId
	return s.db.Save(&roomBooking).Error
}

func (s *RoomBookingService) DeleteRoomBooking(id string) error {
	return s.db.Where("id = ?", id).Delete(&models.RoomBooking{}).Error
}
