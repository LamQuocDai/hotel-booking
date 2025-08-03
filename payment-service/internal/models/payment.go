package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type MethodPayment string

var (
	Cash     MethodPayment = "c"
	Transfer MethodPayment = "t"
)

type Payment struct {
	ID          uuid.UUID     `bson:"_id,omitempty" validate:"required,uuid"`
	BookingId   uuid.UUID     `bson:"booking_id" validate:"required"`
	Method      MethodPayment `bson:"payment_method" validate:"required,oneof=c t"`
	PromotionId uuid.UUID     `bson:"promotion_id, omitempty"`
	SubTotal    int           `bson:"sub_total" validate:"min=0"`
	Discount    int           `bson:"discount" validate:"min=0, max = 100"`
	Tax         int           `bson:"tax" validate:"min=0"`
	Total       int           `bson:"total" validate:"min=0"`
	CreatedAt   time.Time     `bson:"created_at" validate:"required"`
	DeletedAt   *time.Time    `bson:"deleted_at,omitempty"`
}

func (p *Payment) IsValid() error {
	validate := validator.New()
	return validate.Struct(p)
}

func (p *Payment) BeforeCreate() {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	if p.CreatedAt.IsZero() {
		p.CreatedAt = time.Now()
	}
}
