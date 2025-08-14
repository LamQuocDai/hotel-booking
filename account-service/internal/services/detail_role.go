package services

import (
	"my-app/internal/models"

	"gorm.io/gorm"
)

type DetailRoleService struct {
	db *gorm.DB
}

func NewDetailRoleService(db *gorm.DB) *DetailRoleService {
	return &DetailRoleService{db: db}
}

func (s *DetailRoleService) GetAllDetailRoles() ([]models.DetailRole, error) {
	var detailRoles []models.DetailRole
	return detailRoles, s.db.Find(&detailRoles).Error
}

func (s *DetailRoleService) GetDetailRoleByID(id string) (*models.DetailRole, error) {
	var detailRole models.DetailRole
	return &detailRole, s.db.Where("id = ?", id).Preload("Role").Preload("Permission").First(&detailRole).Error
}

func (s *DetailRoleService) CreateDetailRole(detailRole *models.DetailRole) error {
	return s.db.Create(&detailRole).Error
}

func (s *DetailRoleService) UpdatedDetailRole(id string, updatedDetailRole *models.DetailRole) error {
	var detailRole models.DetailRole
	if err := s.db.Where("id = ?", id).First(&detailRole).Error; err != nil {
		return err
	}
	detailRole.RoleId = updatedDetailRole.RoleId
	detailRole.PermissionId = updatedDetailRole.PermissionId
	return s.db.Save(&detailRole).Error
}

func (s *DetailRoleService) DeleteDetailRole(id string) error {
	return s.db.Where("id = ?", id).Delete(&models.DetailRole{}).Error
}
