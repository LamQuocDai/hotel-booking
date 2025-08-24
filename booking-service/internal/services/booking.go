package services

import (
	proto "account-service/proto"
	"booking-service/internal/models"
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookingService struct {
	collection  *mongo.Collection
	redisClient *redis.Client
	proto.UnimplementedBookingServiceServer
}

func NewBookingService(collection *mongo.Collection, redisClient *redis.Client) *BookingService {
	return &BookingService{collection: collection, redisClient: redisClient}
}

func (s *BookingService) CreateBooking(ctx context.Context, req *proto.CreateBookingRequest) (*proto.CreateBookingResponse, error) {
	booking := &models.Booking{
		RoomBookings: make([]uuid.UUID, len(req.RoomIds)),
		UserId:       uuid.MustParse(req.UserId),
		Status:       models.BookingStatus(req.Status),
	}
	for i, id := range req.RoomIds {
		booking.RoomBookings[i] = uuid.MustParse(id)
	}

	booking.BeforeCreate()
	if err := booking.IsValid(); err != nil {
		return &proto.CreateBookingResponse{Error: err.Error()}, nil
	}

	// Acquire Redis locks for all rooms
	var lockKeys []string
	lockKeys = make([]string, len(booking.RoomBookings))
	for i, id := range booking.RoomBookings {
		lockKeys[i] = fmt.Sprintf("lock:room:%s", id.String())
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	locks := make(map[string]*redis.Lock)
	for _, key := range lockKeys {
		lock := s.redisClient.Lock().Key(key).Value(booking.ID.String()).Expiration(10 * time.Second)
		if err := lock.Acquire(ctx); err != nil {
			for _, l := range locks {
				l.Release(ctx)
			}
			return &proto.CreateBookingResponse{Error: fmt.Sprintf("failed to acquire lock for %s: %v", key, err)}, nil
		}
		locks[key] = lock
	}
	defer func() {
		for _, lock := range locks {
			lock.Release(ctx)
		}
	}()

	// Start MongoDB transaction
	session, err := s.collection.Database().Client().StartSession()
	if err != nil {
		return &proto.CreateBookingResponse{Error: err.Error()}, nil
	}
	defer session.EndSession(ctx)

	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		if err := session.StartTransaction(); err != nil {
			return err
		}

		// Check for overlapping bookings
		filter := bson.M{
			"room_bookings": bson.M{"$in": booking.RoomBookings},
			"status":        bson.M{"$in": []models.BookingStatus{Holding, Pending}},
			"deleted_at":    nil,
		}
		var existingBooking models.Booking
		err := s.collection.FindOne(sc, filter).Decode(&existingBooking)
		if err == nil {
			return fmt.Errorf("rooms are already booked in Holding or Pending status")
		} else if err != mongo.ErrNoDocuments {
			return err
		}

		// Insert the booking
		_, err = s.collection.InsertOne(sc, booking)
		if err != nil {
			return err
		}
		return session.CommitTransaction(sc)
	})
	if err != nil {
		session.AbortTransaction(ctx)
		return &proto.CreateBookingResponse{Error: err.Error()}, nil
	}

	return &proto.CreateBookingResponse{Message: "Booking created"}, nil
}
