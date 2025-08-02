package database

import (
	"context"
	"fmt"
	"log"
	"service-service/internal/config"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB(cfg *config.Config) (*mongo.Database, error) {
	if cfg.DBURL == "" {
		return nil, fmt.Errorf("DB_URL is empty")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect cluster
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.DBURL))
	if err != nil {
		return nil, fmt.Errorf("failed to connect db: %v", err)
	}

	// verify the connection
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	//Extract and validate the database name from the URL
	dbName := cfg.DBURL[strings.LastIndex(cfg.DBURL, "/")+1 : strings.Index(cfg.DBURL, "?")]
	if dbName == "" {
		dbName = "service_db" // Default to service_db if not specificed
		log.Printf("No database specific in Url, defaulting to %v", err)
	}
	db := client.Database(dbName)

	return db, nil
}
