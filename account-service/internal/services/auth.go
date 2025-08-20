package services

import (
	"context"
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

	// Bound the query time (helps with remote DB)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := s.db.WithContext(ctx).Where("email = ?", email).First(&acc).Error; err != nil {
		return nil, err
	}
	if err := acc.CheckPassword(password); err != nil {
		return nil, err
	}
	now := time.Now()
	claims := jwt.MapClaims{
		"sub":     acc.ID.String(),
		"email":   acc.Email,
		"role_id": acc.RoleId.String(),
		"iat":     now.Unix(),
		"exp":     now.Add(s.jwtTTL).Unix(),
		"iss":     "account-service",
	}
	// Use HS256 with shared secret (matches your []byte secret)
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString(s.jwtSecret)
	if err != nil {
		return nil, err
	}
	acc.Password = ""
	return &LoginResult{Token: token, Account: &acc}, nil
}
