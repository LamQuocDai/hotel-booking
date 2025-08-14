package models

import (
	"my-app/internal/util"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
	ID        uuid.UUID        `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string           `gorm:"type:varchar(255);not null" validate:"required,max=255"`
	Birthday  *util.CustomTime `gorm:"type:date"`
	Email     string           `gorm:"type:varchar(255);unique;not null" validate:"required,email,max=255"`
	Phone     string           `gorm:"type:varchar(255);unique;not null" validate:"required,max=255"`
	Info      string           `gorm:"type:varchar(255)" validate:"max=255"`
	Vip       VipClass         `gorm:"type:vip_status;default:'nm'" validate:"oneof=nm vp 2vp 3vp b"`
	Password  string           `gorm:"type:varchar(255);not null" validate:"required,min=6,max=255"`
	RoleId    uuid.UUID        `gorm:"type:uuid;not null" validate:"required,uuid"`
	Role      *Role            `gorm:"foreignKey:RoleId;references:ID"`
	CreatedAt time.Time        `gorm:"type:timestamp;default:now()"`
	DeletedAt gorm.DeletedAt   `gorm:"type:timestamp;index"`
}

func (a *Account) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}

	// hash
	if a.Password != "" && !isBcryptHash(a.Password) {
		hashed, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		a.Password = string(hashed)
	}
	return nil
}

func (a *Account) BeforeUpdate(tx *gorm.DB) error {
	if a.Password != "" && !isBcryptHash(a.Password) {
		hashed, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		a.Password = string(hashed)
	}
	return nil
}

func (a *Account) CheckPassword(plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(plain))
}

func isBcryptHash(s string) bool {
	return strings.HasPrefix(s, "$2a$") || strings.HasPrefix(s, "$2b$") || strings.HasPrefix(s, "$2y$")
}
