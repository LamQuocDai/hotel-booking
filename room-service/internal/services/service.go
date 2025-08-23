package services

import (
	"room-service/internal/models"

	"gorm.io/gorm"
)

type ServiceService struct {
	db *gorm.DB
}

func NewServiceService(db *gorm.DB) *ServiceService {
	return &ServiceService{db: db}
}

func (s *ServiceService) GetAllServices() ([]models.Service, error) {
	var services []models.Service
	return services, s.db.Find(&services).Error
}

func (s *ServiceService) GetServiceByID(id string) (*models.Service, error) {
	var service models.Service
	return &service, s.db.Preload("RoomBookings").First(&service, "id = ?", id).Error
}

func (s *ServiceService) CreateService(service *models.Service) error {
	return s.db.Create(service).Error
}

func (s *ServiceService) UpdateService(id string, updatedService *models.Service) error {
	var service models.Service
	if err := s.db.First(&service, "id = ?", id).Error; err != nil {
		return err
	}
	service.Name = updatedService.Name
	service.Description = updatedService.Description
	service.Price = updatedService.Price
	return s.db.Save(&service).Error
}

func (s *ServiceService) DeleteService(id string) error {
	return s.db.Delete(&models.Service{}, "id = ?", id).Error
}
