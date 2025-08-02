package services

import (
	"context"
	"service-service/internal/models"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceService struct {
	collection *mongo.Collection
}

func NewServiceService(db *mongo.Database) *ServiceService {
	return &ServiceService{collection: db.Collection("services")}
}

func (s *ServiceService) GetAllServices() ([]models.Service, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var services []models.Service
	return services, cursor.All(ctx, &services)
}

func (s *ServiceService) GetServiceByID(id string) (*models.Service, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	var service models.Service
	return &service, s.collection.FindOne(ctx, bson.M{"_id": uuid, "delete_at": nil}).Decode(&service)
}

func (s *ServiceService) CreateService(service *models.Service) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	service.BeforeCreate()
	if err := service.IsValid(); err != nil {
		return err
	}
	_, err := s.collection.InsertOne(ctx, service)
	return err
}

func (s *ServiceService) UpdatedService(id string, updatedService *models.Service) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	updatedService.ID = uuid
	if err := updatedService.IsValid(); err != nil {
		return err
	}
	_, err = s.collection.UpdateOne(ctx, bson.M{"_id": uuid}, bson.M{"$set": updatedService})
	return err
}

func (s *ServiceService) DeleteService(id string) error {
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
