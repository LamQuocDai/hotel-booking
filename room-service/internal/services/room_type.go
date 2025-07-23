package services

import (
	"room-service/internal/models"

	"gorm.io/gorm"
)

type RoomTypeService struct {
	db *gorm.DB
}

func NewRoomTypeService(db *gorm.DB) *RoomTypeService {
	return &RoomTypeService{db: db}
}

func (s *RoomTypeService) GetAllRoomTypes() ([]models.RoomType, error) {
	var roomTypes []models.RoomType
	return roomTypes, s.db.Preload("Rooms").Find(roomTypes).Error
}

func (s *RoomTypeService) GetRoomTypeByID(id string) (*models.RoomType, error) {
	var roomType models.RoomType
	return &roomType, s.db.Preload("Rooms").Where("id = ?", id).First(&roomType).Error
}

func (s *RoomTypeService) CreateRoomType(roomType *models.RoomType) error {
	return s.db.Create(roomType).Error
}

func (s *RoomTypeService) UpdatedRoomType(id string, updateRoomType *models.RoomType) error {
	var roomType models.RoomType
	if err := s.db.Where("id = ?", id).First(&roomType).Error; err != nil {
		return err
	}
	roomType.Name = updateRoomType.Name
	roomType.PricePerHour = updateRoomType.PricePerHour
	return s.db.Save(&roomType).Error
}

func (s *RoomTypeService) DeleteRoomType(id string) error {
	return s.db.Where("id = ?", id).Delete(&models.RoomType{}).Error
}
