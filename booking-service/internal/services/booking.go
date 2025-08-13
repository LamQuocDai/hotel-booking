package services

import (
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookingService struct {
	collection  *mongo.Collection
	redisClient *redis.Client
}
