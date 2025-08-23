package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	ID           uuid.UUID       `gorm:"type:uuid;primaryKey"`
	Name         string          `gorm:"type:string;size:255;not null" validate:"required"`
	Description  string          `gorm:"type:string;size:255"`
	Price        int             `gorm:"type:int;not null" validate:"required,min=0"`
	RoomBookings []RoomBooking   `gorm:"foreignKey:ServiceId;references:ID"`
	CreatedAt    time.Time       `gorm:"type:timestamp;default:now()"`
	DeletedAt    *gorm.DeletedAt `gorm:"type:timestamp;index"`
}

func (s *Service) BeforeCreate(tx *gorm.DB) error {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return nil
}
