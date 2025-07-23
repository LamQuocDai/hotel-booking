package database

import (
	"fmt"
	"room-service/internal/config"
	"room-service/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	if cfg.DBURL == "" {
		return nil, fmt.Errorf("DB_URL is not set")
	}

	// Step 1: Temporary connection for setup
	tempDB, err := gorm.Open(postgres.Open(cfg.DBURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("connect db failed: %v", err)
	}

	// Step 2: Ensure room schema and extension exist
	if err := tempDB.Exec("CREATE SCHEMA IF NOT EXISTS room").Error; err != nil {
		return nil, fmt.Errorf("failed to create room schema: %v", err)
	}

	if err := tempDB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		return nil, fmt.Errorf("failed to enable uuid extension: %v", err)
	}

	if err := tempDB.Exec(`DO $$ BEGIN 
		CREATE TYPE room_status AS ENUM ('occ','ooo','cl'); 
		EXCEPTION WHEN duplicate_object THEN NULL; 
	END $$;`).Error; err != nil {
		return nil, fmt.Errorf("failed to create room_status enum: %v", err)
	}

	// Step 3: Reconnect with search_path set in DSN
	dsn := cfg.DBURL + "?search_path=room,public"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("connect db with search_path failed: %v", err)
	}

	// Step 4: Run migrations (in "room" schema automatically)
	if err := db.AutoMigrate(
		&models.Room{},
		&models.Location{},
		&models.RoomImage{},
		&models.RoomType{},
		&models.Review{},
	); err != nil {
		return nil, fmt.Errorf("migration failed: %v", err)
	}

	return db, nil
}
