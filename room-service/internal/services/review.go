package services

import (
	"room-service/internal/models"

	"gorm.io/gorm"
)

type ReviewService struct {
	db *gorm.DB
}

func NewReviewService(db *gorm.DB) *ReviewService {
	return &ReviewService{db: db}
}

func (s *ReviewService) GetAllReviews() ([]models.Review, error) {
	var reviews []models.Review
	return reviews, s.db.Preload("Room").Find(&reviews).Error
}

func (s *ReviewService) GetReviewByID(id string) (*models.Review, error) {
	var review models.Review
	return &review, s.db.Preload("Room").Find(&review).Error
}

func (s *ReviewService) CreateReview(review *models.Review) error {
	return s.db.Create(&review).Error
}

func (s *ReviewService) UpdatedReview(id string, updatedReview *models.Review) error {
	var review models.Review
	if err := s.db.Where("id = ?", id).First(&review).Error; err != nil {
		return err
	}
	review.Rating = updatedReview.Rating
	review.Comment = updatedReview.Comment
	return s.db.Save(&review).Error
}

func (s *ReviewService) DeleteReview(id string) error {
	return s.db.Where("id = ?", id).Delete(&models.Review{}).Error
}
