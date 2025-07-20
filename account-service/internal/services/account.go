package services

import (
	"my-app/internal/models"

	"gorm.io/gorm"
)

type AccountService struct {
	db *gorm.DB
}

func NewAccountService(db *gorm.DB) *AccountService {
	return &AccountService{db: db}
}

func (s *AccountService) GetAllCounts() ([]models.Account, error) {
	var accounts []models.Account
	if err := s.db.Preload("Role").Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

func (s *AccountService) GetAccountByID(id string) (*models.Account, error) {
	var account models.Account
	if err := s.db.Preload("Role").Where("id = ?", id).First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (s *AccountService) CreateAccount(account *models.Account) error {
	return s.db.Create(account).Error
}

func (s *AccountService) UpdateAccount(id string, updatedAccount *models.Account) error {
	var account models.Account
	if err := s.db.Where("id = ?", id).First(&account).Error; err != nil {
		return err
	}
	account.Name = updatedAccount.Name
	account.Birthday = updatedAccount.Birthday
	account.Info = updatedAccount.Info
	account.Vip = updatedAccount.Vip
	account.Password = updatedAccount.Password
	account.RoleId = updatedAccount.RoleId
	return s.db.Save(&account).Error
}

func (s *AccountService) DeleteAccount(id string) error {
	return s.db.Where("id = ?", id).Delete(&models.Account{}).Error
}
