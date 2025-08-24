package models

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Promotion struct {
	ID          uuid.UUID  `bson:"_id,omitempty" json:"id"`
	Code        string     `bson:"code" json:"code" validate:"required"`
	Description string     `bson:"description" json:"description"`
	Discount    int        `bson:"discount" json:"discount" validate:"required,min=0,max=100"`
	StartDay    int32      `bson:"start_date" json:"startDate" validate:"required"`
	EndDay      int32      `bson:"end_date" json:"endDate" validate:"required"`
	CreatedAt   time.Time  `bson:"created_at" json:"createdAt"`
	DeletedAt   *time.Time `bson:"deleted_at,omitempty" json:"deletedAt,omitempty"`
}

func (p *Promotion) BeforeCreate() {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	if p.CreatedAt.IsZero() {
		p.CreatedAt = time.Now()
	}
}

func (p *Promotion) IsValid() error {
	v := validator.New()
	return v.Struct(p)
}

func (p *Promotion) Validate() error {
	if p.Code == "" {
		return errors.New("code required")
	}
	if p.Discount < 0 || p.Discount > 100 {
		return errors.New("discount out of range")
	}
	if p.StartDay == 0 || p.EndDay == 0 {
		return errors.New("startDay/endDay required")
	}
	if p.EndDay < p.StartDay {
		return errors.New("endDay before startDay")
	}
	return nil
}
