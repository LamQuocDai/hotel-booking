package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Permission struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string         `gorm:"type:varchar(255);not null" validate:"required"`
	Access      string         `gorm:"type:varchar(255);not null" validate:"required"`
	DetailRoles []DetailRole   `gorm:"foreignKey:PermissionId;references:ID"`
	CreatedAt   time.Time      `gorm:"type:timestamp;default:now()"`
	DeletedAt   gorm.DeletedAt `gorm:"type:timestamp;index"`
}

func (p *Permission) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}
