package services

import (
	"room-service/internal/models"

	"gorm.io/gorm"
)

type RoomImageService struct {
	db *gorm.DB
}

func NewRoomImageService(db *gorm.DB) *RoomImageService {
	return &RoomImageService{db: db}
}

func (s *RoomImageService) GetAllRoomImages() ([]models.RoomImage, error) {
	var roomImages []models.RoomImage
	return roomImages, s.db.Preload("Room").Find(&roomImages).Error
}

func (s *RoomImageService) GetRoomImageByID(id string) (*models.RoomImage, error) {
	var roomImage models.RoomImage
	return &roomImage, s.db.Preload("Room").Where("id = ?", id).First(&roomImage).Error
}

func (s *RoomImageService) CreateRoomImage(roomImage *models.RoomImage) error {
	return s.db.Create(&roomImage).Error
}

func (s *RoomImageService) UpdatedRoomImage(id string, updatedRoomImage *models.RoomImage) error {
	var roomImage models.RoomImage
	if err := s.db.Where("id = ?", id).First(&roomImage).Error; err != nil {
		return err
	}
	roomImage.RoomId = updatedRoomImage.RoomId
	roomImage.ImageURL = updatedRoomImage.ImageURL
	return s.db.Save(&roomImage).Error
}

func (s *RoomImageService) DeleteRoomImage(id string) error {
	return s.db.Where("id = ?", id).Delete(&models.RoomImage{}).Error
}
