package services

import (
	"context"
	"service-service/internal/models"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceBookingService struct {
	collection *mongo.Collection
}

func NewServiceBookingService(db *mongo.Database) *ServiceBookingService {
	return &ServiceBookingService{collection: db.Collection("service_bookings")}
}

func (s *ServiceBookingService) GetAllServiceBookings() ([]models.ServiceBooking, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var serviceBookings []models.ServiceBooking
	return serviceBookings, cursor.All(ctx, &serviceBookings)
}

func (s *ServiceBookingService) GetServiceBookingByID(id string) (*models.ServiceBooking, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	var serviceBooking models.ServiceBooking
	return &serviceBooking, s.collection.FindOne(ctx, bson.M{"_id": uuid}).Decode(&serviceBooking)
}

func (s *ServiceBookingService) CreateServiceBooking(serviceBooking *models.ServiceBooking) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	serviceBooking.BeforeCreate()
	if err := serviceBooking.IsValid(); err != nil {
		return err
	}
	_, err := s.collection.InsertOne(ctx, &serviceBooking)
	return err
}

func (s *ServiceBookingService) UpdatedServiceBooking(id string, updatedServiceBooking *models.ServiceBooking) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	updatedServiceBooking.ID = uuid
	if err := updatedServiceBooking.IsValid(); err != nil {
		return err
	}
	_, err = s.collection.UpdateOne(ctx, bson.M{"_id": uuid}, bson.M{"$set": updatedServiceBooking})
	return err
}

func (s *ServiceBookingService) DeleteServiceBooking(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil
	}
	_, err = s.collection.DeleteOne(ctx, bson.M{"_id": uuid})
	return err
}
