package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Promotion struct {
	ID          uuid.UUID  `bson:"_id, omitemtpy" validate:"required,uuuid"`
	Code        string     `bson:"code" validate:"required"`
	Description string     `bson:"description"`
	Discount    int        `bson:"discount" validate:"required, min=0, max=100"`
	StartDate   time.Time  `bson:"start_date" validate:"required"`
	EndDate     time.Time  `bson:"end_date" validate:"required"`
	CreatedAt   time.Time  `bson:"created_at" validate:"required"`
	DeletedAt   *time.Time `bson:"deleted_at, omitempty"`
}

func (p *Promotion) IsValid() error {
	validate := validator.New()
	return validate.Struct(p)
}

func (p *Promotion) BeforeCreate() {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	if p.CreatedAt.IsZero() {
		p.CreatedAt = time.Now()
	}
}
