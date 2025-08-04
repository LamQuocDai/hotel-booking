package services

import (
	"context"
	"payment-service/internal/models"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PromotionService struct {
	collection *mongo.Collection
}

func NewPromotionService(db *mongo.Database) *PromotionService {
	return &PromotionService{collection: db.Collection("promotions")}
}

func (s *PromotionService) GetAllPromotions() ([]models.Promotion, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var promotions []models.Promotion
	return promotions, cursor.All(ctx, &promotions)
}

func (s *PromotionService) GetPromotionByID(id string) (*models.Promotion, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	var promotion models.Promotion
	return &promotion, s.collection.FindOne(ctx, bson.M{"_id": uuid, "deleted_at": nil}).Decode(&promotion)
}

func (s *PromotionService) CreatePromotion(promotion *models.Promotion) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	promotion.BeforeCreate()
	if err := promotion.IsValid(); err != nil {
		return err
	}
	_, err := s.collection.InsertOne(ctx, &promotion)
	return err
}

func (s *PromotionService) UpdatedPromotion(id string, updatePromotion *models.Promotion) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	if err := updatePromotion.IsValid(); err != nil {
		return err
	}
	_, err = s.collection.UpdateOne(ctx, bson.M{"_id": uuid}, bson.M{"$set": updatePromotion})
	return err
}

func (s *PromotionService) DeletePromotion(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil
	}
	now := time.Now()
	_, err = s.collection.UpdateOne(ctx, bson.M{"_id": uuid}, bson.M{"$set": bson.M{"deleted_at": &now}})
	return err
}
