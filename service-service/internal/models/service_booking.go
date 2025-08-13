package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ServiceBooking struct {
	ID        uuid.UUID `bson:"_id,omitempty" validate:"required,uuid"`
	ServiceId uuid.UUID `bson:"service_id" validate:"required,uuid"`
	BookingId uuid.UUID `bson:"booking_id" validate:"required,uuid"`
	Quantity  int       `bson:"quantity" validate:"required,min=0"`
	Total     int       `bson:"total" validate:"requried,min=0"`
}

func (s *ServiceBooking) IsValid() error {
	validate := validator.New()
	return validate.Struct(s)
}

func (s *ServiceBooking) BeforeCreate() {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
}
