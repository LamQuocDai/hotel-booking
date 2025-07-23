package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Location struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string         `gorm:"type:varchar(255);not null;unique" validate:"required"`
	Address     string         `gorm:"type:varchar(255);not null;unique" validate:"required"`
	Description string         `gorm:"type:varchar(255);not null"`
	Rooms       []Room         `gorm:"foreignKey:LocationId;references:ID"`
	CreatedAt   time.Time      `gorm:"type:timestamp;default:now()"`
	DeletedAt   gorm.DeletedAt `gorm:"type:timestamp;index"`
}

func (l *Location) BeforeCreate(tx *gorm.DB) error {
	if l.ID == uuid.Nil {
		l.ID = uuid.New()
	}
	return nil
}
