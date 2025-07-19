package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VipClass string

var (
	Normal VipClass = "nm"
	Vip    VipClass = "vp"
	Vvip   VipClass = "2vp"
	Vvvip  VipClass = "3vp"
	Banned VipClass = "b"
)

type Account struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string         `gorm:"type:varchar(255);not null" validate:"required,max=255"`
	Birthday  *time.Time     `gorm:"type:date"`
	Email     string         `gorm:"type:varchar(255);unique;not null" validate:"required,email,max=255"`
	Phone     string         `gorm:"type:varchar(255);unique;not null" validate:"required,max=255"`
	Info      string         `gorm:"type:varchar(255)" validate:"max=255"`
	Vip       VipClass       `gorm:"type:vip_status;default:'nm'" validate:"required,oneof=nm vp 2vp 3vp b"`
	Password  string         `gorm:"type:varchar(255);not null" validate:"required,min=6,max=255"`
	RoleId    uuid.UUID      `gorm:"type:uuid;not null" validate:"required,uuid"`
	Role      Role           `gorm:"foreignKey:RoleId;references:ID"`
	CreatedAt time.Time      `gorm:"type:timestamp;default:now()"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp;index"`
}

func (a *Account) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	return nil
}
