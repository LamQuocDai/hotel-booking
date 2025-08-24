package database

import (
	"context"
	"fmt"
	"payment-service/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitDB(cfg *config.Config) (*mongo.Database, error) {
	if cfg.DBURL == "" {
		return nil, fmt.Errorf("DB_URL is emppty")
	}
	if cfg.DBName == "" {
		cfg.DBName = "payment_db"
	}

	clientOpts := options.Client().
		ApplyURI(cfg.DBURL).
		SetServerSelectionTimeout(15 * time.Second)

	connectCtx, connectCancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer connectCancel()

	client, err := mongo.Connect(connectCtx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect db: %v", err)
	}

	pingCtx, pingCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer pingCancel()
	if err := client.Ping(pingCtx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return client.Database(cfg.DBName), nil
}
