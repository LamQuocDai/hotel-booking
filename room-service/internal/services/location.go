package services

import (
	"room-service/internal/models"

	"gorm.io/gorm"
)

type LocationService struct {
	db *gorm.DB
}

func NewLocationService(db *gorm.DB) *LocationService {
	return &LocationService{db: db}
}

func (s *LocationService) GetAllLocations() ([]models.Location, error) {
	var locations []models.Location
	return locations, s.db.Preload("Rooms").Find(&locations).Error
}

func (s *LocationService) GetLocationByID(id string) (*models.Location, error) {
	var location models.Location
	return &location, s.db.Preload("Rooms").Where("id = ?", id).First(&location).Error
}

func (s *LocationService) CreateLocation(location *models.Location) error {
	return s.db.Create(location).Error
}

func (s *LocationService) UpdatedLocation(id string, updatedLocation *models.Location) error {
	var location models.Location
	// Check ID existed
	if err := s.db.Where("id = ?", id).First(&location).Error; err != nil {
		return err
	}
	location.Name = updatedLocation.Name
	location.Address = updatedLocation.Address
	location.Description = updatedLocation.Description
	return s.db.Save(&location).Error
}

func (s *LocationService) DeleteLocation(id string) error {
	return s.db.Where("id = ?", id).Delete(&models.Location{}).Error
}
