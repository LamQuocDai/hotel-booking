package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoomType struct {
	ID           uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name         string         `gorm:"type:varchar(255);not null;unique" validate:"required"`
	PricePerHour int            `gorm:"type:integer;not null" validate:"required"`
	Rooms        []Room         `gorm:"type:foreignKey:RoomTypeId"`
	CreatedAt    time.Time      `gorm:"type:timestamp;default:now()"`
	DeletedAt    gorm.DeletedAt `gorm:"type:timestamp;index"`
}

func (rt *RoomType) BeforeCreate(tx *gorm.DB) error {
	if rt.ID == uuid.Nil {
		rt.ID = uuid.New()
	}
	return nil
}
