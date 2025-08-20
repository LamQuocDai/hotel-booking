package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey"`
	AccountId uuid.UUID      `gorm:"type:uuid;not null" validate:"required"`
	RoomId    uuid.UUID      `gorm:"type:uuid;not null" validate:"required"`
	Room      *Room          `gorm:"foreignKey:RoomId;references:ID"`
	Rating    int            `gorm:"type:smallint" validate:"min=0,max=5"`
	Comment   string         `gorm:"type:varchar(255)"`
	CreatedAt time.Time      `gorm:"type:timestamp;default:now()"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp;index"`
}

func (r *Review) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return nil
}
