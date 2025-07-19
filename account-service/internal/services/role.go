package services

import (
	"my-app/internal/models"

	"gorm.io/gorm"
)

type RoleService struct {
	db *gorm.DB
}

func NewRoleService(db *gorm.DB) *RoleService {
	return &RoleService{db: db}
}

func (s *RoleService) GetAllRoles() ([]models.Role, error) {
	var roles []models.Role
	if err := s.db.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (s *RoleService) GetRoleByID(id string) (*models.Role, error) {
	var role models.Role
	if err := s.db.Where("id = ?", id).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (s *RoleService) CreateRole(role *models.Role) error {
	return s.db.Create(role).Error
}

func (s *RoleService) UpdateRole(id string, updateRole *models.Role) error {
	var role models.Role
	if err := s.db.Where("id = ?", id).First(&role).Error; err != nil {
		return err
	}
	role.Name = updateRole.Name
	role.Description = updateRole.Description
	return s.db.Save(&role).Error
}

func (s *RoleService) DeleteRole(id string) error {
	return s.db.Where("id = ?", id).Delete(&models.Role{}).Error
}
