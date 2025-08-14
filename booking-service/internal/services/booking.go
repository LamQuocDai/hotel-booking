package services

import (
	"booking-service/internal/models"
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookingService struct {
	collection  *mongo.Collection
	redisClient *redis.Client
}

func NewBookingService(db *mongo.Database, rdb *redis.Client) *BookingService {
	return &BookingService{collection: db.Collection("bookings"), redisClient: rdb}
}

func (s *BookingService) acquireLock(ctx context.Context, resourceID string) (bool, error) {
	lockKey := "lock:" + resourceID
	return s.redisClient.SetNX(ctx, lockKey, "locked", 10*time.Second).Result()
}

func (s *BookingService) releaseLock(ctx context.Context, resourceID string) error {
	lockKey := "lock:" + resourceID
	return s.redisClient.Del(ctx, lockKey).Err()
}

func (s *BookingService) CreateBooking(booking *models.Booking) error {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// // Determine resource ID based on type
	// var resourceID string
	// if booking.Type == models.BookingTypeService {
	// 	if booking.ServiceID == nil {
	// 		return fmt.Errorf("service_id is required for service booking")
	// 	}
	// 	resourceID = booking.ServiceID.String()
	// } else if booking.Type == models.BookingTypeRoom {
	// 	if booking.RoomID == nil {
	// 		return fmt.Errorf("room_id is required for room booking")
	// 	}
	// 	resourceID = booking.RoomID.String()
	// } else {
	// 	return fmt.Errorf("invalid booking type")
	// }

	// // Acquire a lock to prevent concurrent bookings
	// acquired, err := s.acquireLock(ctx, resourceID)
	// if err != nil || !acquired {
	// 	return fmt.Errorf("failed to acquire lock: %v", err)
	// }
	// defer s.releaseLock(ctx, resourceID)

	// // Check for overlapping bookings
	// var existingBooking models.Booking
	// err = s.collection.FindOne(ctx, bson.M{
	// 	"type":      booking.Type,
	// 	"service_id": booking.ServiceID,
	// 	"room_id":    booking.RoomID,
	// 	"start_time": bson.M{"$lte": booking.EndTime},
	// 	"end_time":   bson.M{"$gte": booking.StartTime},
	// 	"deleted_at": nil,
	// }).Decode(&existingBooking)
	// if err == nil {
	// 	return fmt.Errorf("slot already booked")
	// } else if err != mongo.ErrNoDocuments {
	// 	return err
	// }

	// // Create the booking
	// booking.BeforeCreate()
	// if err := booking.IsValid(); err != nil {
	// 	return err
	// }
	// _, err = s.collection.InsertOne(ctx, booking)

	return nil
}

func (s *BookingService) GetBookingByID(id string) (*models.Booking, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	var booking models.Booking
	return &booking, s.collection.FindOne(ctx, bson.M{"_id": uuid, "deleted_at": nil}).Decode(&booking)
}

func (s *BookingService) UpdateBookingStatus(id string, status models.BookingStatus) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	if err := status.IsValidStatus(); err != nil {
		return err
	}
	_, err = s.collection.UpdateOne(ctx, bson.M{"_id": uuid, "deleted_at": nil}, bson.M{"$set": bson.M{"status": status}})
	return err
}

func (s *BookingService) DeleteBooking(id string) error {
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
