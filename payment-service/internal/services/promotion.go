package services

import (
	"context"
	"errors"
	"fmt"
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

	// match not soft-deleted (field missing OR null)
	filter := bson.M{"$or": []bson.M{
		{"deleted_at": bson.M{"$exists": false}},
		{"deleted_at": nil},
	}}

	cursor, err := s.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var promotions []models.Promotion
	if err := cursor.All(ctx, &promotions); err != nil {
		return nil, err
	}
	return promotions, nil
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
	fmt.Println(1)
	if err := promotion.IsValid(); err != nil {
		return err
	}
	if err := promotion.Validate(); err != nil {
		return err
	}
	fmt.Println(2)
	_, err := s.collection.InsertOne(ctx, &promotion)
	fmt.Println(3)
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
	if err := updatePromotion.Validate(); err != nil {
		return err
	}

	// Only update mutable fields
	update := bson.M{
		"code":        updatePromotion.Code,
		"description": updatePromotion.Description,
		"discount":    updatePromotion.Discount,
		"start_day":   updatePromotion.StartDay,
		"end_day":     updatePromotion.EndDay,
	}
	_, err = s.collection.UpdateOne(ctx, bson.M{"_id": uuid}, bson.M{"$set": update})
	return err
}

func (s *PromotionService) DeletePromotion(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil
	}
	now := time.Now().UTC()
	res, err := s.collection.UpdateOne(ctx, bson.M{"_id": uuid}, bson.M{
		"$set": bson.M{"deleted_at": &now},
	})
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return errors.New("promotion not found")
	}
	return err
}
