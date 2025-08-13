package models

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type BookingStatus string

var (
	Holding   BookingStatus = "h"
	Pending   BookingStatus = "p"
	Confirmed BookingStatus = "cf"
	CheckIn   BookingStatus = "ci"
	CheckOut  BookingStatus = "co"
	Cancelled BookingStatus = "cl"
)

func (bs BookingStatus) IsValidStatus() error {
	validStatuses := []BookingStatus{Holding, Pending, Confirmed, CheckIn, CheckOut, Cancelled}
	for _, valid := range validStatuses {
		if bs == valid {
			return nil
		}
	}
	return fmt.Errorf("invalid booking status: %s", bs)
}

type Booking struct {
	ID              uuid.UUID     `bson:"_id,omitempty" validate:"required,uuid"`
	UserId          uuid.UUID     `bson:"user_id" validate:"required,uuid"`
	CheckIn         time.Time     `bson:"check_in" validate:"requried"`
	CheckOut        time.Time     `bson:"check_out" validate:"required,gtfield=CheckIn"`
	Status          BookingStatus `bson:"status" validate:"required"`
	ServiceBookings []uuid.UUID   `bson:"service_bookings" validate:"dive"`
	RoomBookings    []uuid.UUID   `bson:"room_bookings" validate:"dive"`
	CreatedAt       time.Time     `bson:"created_at" validate:"required"`
	DeletedAt       *time.Time    `bson:"deleted_at,omitempty"`
}

func (b *Booking) IsValid() error {
	validate := validator.New()
	if err := validate.Struct(b); err != nil {
		return err
	}
	if err := b.Status.IsValidStatus(); err != nil {
		return err
	}
	return nil
}

func (b *Booking) BeforeCreate() {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	if b.CreatedAt.IsZero() {
		b.CreatedAt = time.Now()
	}
	if b.Status == "" {
		b.Status = Pending
	}
}
