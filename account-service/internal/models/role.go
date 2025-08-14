package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string         `gorm:"type:varchar(255);not null;unique" validate:"required,max=255"`
	Description string         `gorm:"type:varchar(255)" validate:"max=255"`
	Accounts    []Account      `gorm:"foreignKey:RoleId;references:ID"`
	DetailRoles []DetailRole   `gorm:"foreignKey:RoleId;references:ID"`
	CreatedAt   time.Time      `gorm:"type:timestamp;default:now()"`
	DeletedAt   gorm.DeletedAt `gorm:"type:timestamp;index"`
}

func (r *Role) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return nil
}
