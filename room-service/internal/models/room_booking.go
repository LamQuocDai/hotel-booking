package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoomBooking struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey"`
	BookingId uuid.UUID  `gorm:"type:uuid;not null"`
	RoomId    uuid.UUID  `gorm:"type:uuid;not null"`
	Room      *Room      `gorm:"foreignKey:RoomId;references:ID"`
	ServiceId *uuid.UUID `gorm:"type:uuid"`
	Service   *Service   `gorm:"foreignKey:ServiceId;references:ID"`
	CheckIn   *time.Time `gorm:"type:timestamp" validate:"required"`
	CheckOut  *time.Time `gorm:"type:timestamp" validate:"required,gtfield=CheckIn"`
	Price     int        `gorm:"type:int;not null" validate:"required,min=0"`
}

func (rb *RoomBooking) BeforeCreate(tx *gorm.DB) error {
	if rb.ID == uuid.Nil {
		rb.ID = uuid.New()
	}
	return nil
}
