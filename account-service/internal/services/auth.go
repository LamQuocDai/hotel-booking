package services

import (
	"my-app/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AuthService struct {
	db        *gorm.DB
	jwtSecret []byte
	jwtTTL    time.Duration
}

func NewAuthService(db *gorm.DB, secret string, ttl time.Duration) *AuthService {
	return &AuthService{db: db, jwtSecret: []byte(secret), jwtTTL: ttl}
}

type LoginResult struct {
	Token   string
	Account *models.Account
}

func (s *AuthService) Login(email, password string) (*LoginResult, error) {
	var acc models.Account
	if err := s.db.Where("email = !", email).First(&acc).Error; err != nil {
		return nil, err
	}
	if err := acc.CheckPassword(password); err != nil {
		return nil, err
	}
	now := time.Now()
	claims := jwt.MapClaims{
		"sub":     acc.ID.String(),
		"role_id": acc.RoleId.String(),
		"iat":     now.Unix(),
		"exp":     now.Add(s.jwtTTL).Unix(),
		"iss":     "account-service",
	}
	t := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token, err := t.SignedString(s.jwtSecret)
	if err != nil {
		return nil, err
	}
	acc.Password = ""
	return &LoginResult{Token: token, Account: &acc}, nil
}
