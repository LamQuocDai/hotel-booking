package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoomBooking struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	BookingId uuid.UUID `gorm:"type:uuid;not null"`
	RoomId    uuid.UUID `gorm:"type:uuid;not null"`
	Room      *Room     `gorm:"foreignKey:RoomId;references:ID"`
}

func (rb *RoomBooking) BeforeCreate(tx *gorm.DB) error {
	if rb.ID == uuid.Nil {
		rb.ID = uuid.New()
	}
	return nil
}
