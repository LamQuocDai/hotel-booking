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

	db, err := gorm.Open(postgres.Open(cfg.DBURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("conntct db failed: %v", err)
	}

	// Enable uuid extension
	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
		return nil, fmt.Errorf("enable uuid failed: %v", err)
	}

	// Enum room status
	if err := db.Exec(`DO $$ BEGIN
	CREATE TYPE room_status AS ENUM ('ooc','ooo','cl');
	EXCEPTION WHEN duplicate_object THEN NULL;
	END $$;`).Error; err != nil {
		return nil, fmt.Errorf("failed to create room_status enum: %v", err)
	}

	// mirgate
	if err := db.AutoMigrate(&models.Room{}, &models.Location{}, &models.RoomImage{}, &models.RoomType{}); err != nil {
		return nil, fmt.Errorf("migrate failed: %v", err)
	}

	return db, nil
}
