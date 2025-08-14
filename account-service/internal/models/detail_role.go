package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DetailRole struct {
	ID           uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	RoleId       uuid.UUID   `gorm:"type:uuid;not null"`
	Role         *Role       `gorm:"foreignKey:RoleId;references:ID"`
	PermissionId uuid.UUID   `gorm:"type:uuid;not null"`
	Permission   *Permission `gorm:"foreignKey:PermissionId;references:ID"`
}

func (dr *DetailRole) BeforeCreate(tx *gorm.DB) error {
	if dr.ID == uuid.Nil {
		dr.ID = uuid.New()
	}
	return nil
}
