package services

import (
	"my-app/internal/models"

	"gorm.io/gorm"
)

type PermissionService struct {
	db *gorm.DB
}

func NewPermissionService(db *gorm.DB) *PermissionService {
	return &PermissionService{db: db}
}

func (s *PermissionService) GetAllPermissions() ([]models.Permission, error) {
	var permissions []models.Permission
	return permissions, s.db.Find(&permissions).Error
}

func (s *PermissionService) GetPermissionByID(id string) (*models.Permission, error) {
	var permision models.Permission
	return &permision, s.db.Where("id = ?", id).Preload("DetailRoles").First(&permision).Error
}

func (s *PermissionService) CreatePermission(permission *models.Permission) error {
	return s.db.Create(&permission).Error
}

func (s *PermissionService) UpdatedPermission(id string, updatedPermission *models.Permission) error {
	var permission models.Permission
	if err := s.db.Where("id = ?", id).First(&permission).Error; err != nil {
		return err
	}
	permission.Name = updatedPermission.Name
	permission.Access = updatedPermission.Access
	return s.db.Save(&permission).Error
}

func (s *PermissionService) DeletePermission(id string) error {
	return s.db.Where("id = ?", id).Delete(&models.Permission{}).Error
}
