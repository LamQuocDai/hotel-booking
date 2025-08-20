package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoomImage struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	RoomId   uuid.UUID `gorm:"type:uuid;not null"`
	Room     *Room     `gorm:"foreignKey:RoomId;references:ID"`
	ImageURL string    `gorm:"type:varchar(255);not null"`
}

func (ri *RoomImage) BeforeCreate(tx *gorm.DB) error {
	if ri.ID == uuid.Nil {
		ri.ID = uuid.New()
	}
	return nil
}
