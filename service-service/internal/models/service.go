package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Service struct {
	ID              uuid.UUID        `bson:"_id,omitempty" validate:"required,uuid"`
	Name            string           `bson:"name" validate:"required,max=255"`
	Description     string           `bson:"description" validate:"max=255"`
	Price           int              `bson:"price" validate:"required,min=0"`
	ServiceBookings []ServiceBooking `bson:"service_bookings" validate:"dive"`
	CreatedAt       time.Time        `bson:"created_at" validate:"required"`
	DeletedAt       *time.Time       `bson:"deleted_at,omitempty"`
}

func (s *Service) IsValid() error {
	validate := validator.New()
	return validate.Struct(s)
}

func (s *Service) BeforeCreate() {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now()
	}
}
