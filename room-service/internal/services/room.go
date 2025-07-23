package services

import (
	"room-service/internal/models"

	"gorm.io/gorm"
)

type RoomService struct {
	db *gorm.DB
}

func NewRoomService(db *gorm.DB) *RoomService {
	return &RoomService{db: db}
}

func (s *RoomService) GetAllRooms() ([]models.Room, error) {
	var rooms []models.Room
	return rooms, s.db.Preload("Location").Preload("RoomType").Preload("RoomImages").Preload("Reviews").Find(&rooms).Error
}

func (s *RoomService) GetRoomByID(id string) (*models.Room, error) {
	var room models.Room
	return &room, s.db.Preload("Location").Preload("RoomType").Preload("RoomImages").Preload("Reviews").Where("id = ?", id).First(&room).Error
}

func (s *RoomService) CreateRoom(room *models.Room) error {
	return s.db.Create(&room).Error
}

func (s *RoomService) UpdatedRoom(id string, updatedRoom *models.Room) error {
	var room models.Room
	if err := s.db.Where("id = ?", id).First(&room).Error; err != nil {
		return err
	}
	room.Name = updatedRoom.Name
	room.LocationId = updatedRoom.LocationId
	room.RoomTypeId = updatedRoom.RoomTypeId
	room.Status = updatedRoom.Status
	return s.db.Save(&room).Error
}

func (s *RoomService) DeleteRoom(id string) error {
	return s.db.Where("id = ?", id).Delete(&models.Room{}).Error
}
