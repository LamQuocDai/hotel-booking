package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoomStatus string

var (
	Occupied   RoomStatus = "occ"
	OutOfOrder RoomStatus = "ooo"
	Clean      RoomStatus = "cl"
)

type Room struct {
	ID           uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name         string         `gorm:"type:varchar(255);not null;unique" validate:"required"`
	LocationId   uuid.UUID      `gorm:"type:uuid;not null" validate:"required"`
	Location     *Location      `gorm:"foreignKey:LocationId;references:ID"`
	RoomTypeId   uuid.UUID      `gorm:"type:uuid;not null" validate:"required"`
	RoomType     *RoomType      `gorm:"foreignKey:RoomTypeId;references:ID"`
	Status       RoomStatus     `gorm:"type:room_status;default:'ooo'" validate:"oneof=occ ooo cl"`
	RoomImages   []RoomImage    `gorm:"foreignKey:RoomId;references:ID"`
	Reviews      []Review       `gorm:"foreignKey:RoomId;references:ID"`
	RoomBookings []RoomBooking  `gorm:"foreignKey:RoomId;references:ID"`
	CreatedAt    time.Time      `gorm:"type:timestamp;default:now()"`
	DeletedAt    gorm.DeletedAt `gorm:"type:timestamp;index"`
}

func (r *Room) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return nil
}
