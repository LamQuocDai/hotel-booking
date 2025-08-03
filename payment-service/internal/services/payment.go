package services

import (
	"context"
	"payment-service/internal/models"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PaymentService struct {
	collection *mongo.Collection
}

func NewPaymentService(db *mongo.Database) *PaymentService {
	return &PaymentService{collection: db.Collection("payments")}
}

func (s *PaymentService) GetAllPayments() ([]models.Payment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var services []models.Payment
	return services, cursor.All(ctx, &services)
}

func (s *PaymentService) GetPaymentByID(id string) (*models.Payment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	var service models.Payment
	return &service, s.collection.FindOne(ctx, bson.M{"_id": uuid, "deleted_at": nil}).Decode(&service)
}

func (s *PaymentService) CreatePayment(payment *models.Payment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	payment.BeforeCreate()
	if err := payment.IsValid(); err != nil {
		return err
	}
	_, err := s.collection.InsertOne(ctx, payment)
	return err
}

func (s *PaymentService) UpdatedPayment(id string, updatedPayment *models.Payment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	updatedPayment.ID = uuid
	if err := updatedPayment.IsValid(); err != nil {
		return err
	}
	_, err = s.collection.UpdateOne(ctx, bson.M{"_id": uuid}, bson.M{"$set": updatedPayment})
	return err
}

func (s *PaymentService) DeletePayment(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	now := time.Now()
	_, err = s.collection.UpdateOne(ctx, bson.M{"_id": uuid}, bson.M{"$set": bson.M{"deleted_at": &now}})
	return err
}
