package database

import (
	"fmt"

	"room-service/internal/config"
	"room-service/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	if cfg.DBURL == "" {
		return nil, fmt.Errorf("DB_URL is not set")
	}

	// Setup: ensure schema and enum exist
	setupDB, err := gorm.Open(postgres.Open(cfg.DBURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("connect db failed: %v", err)
	}
	if err := setupDB.Exec(`CREATE SCHEMA IF NOT EXISTS room;`).Error; err != nil {
		return nil, fmt.Errorf("failed to create room schema: %v", err)
	}
	if err := setupDB.Exec(`DO $$ BEGIN
        CREATE TYPE room.room_status AS ENUM ('occ','ooo','cl');
    EXCEPTION WHEN duplicate_object THEN NULL;
    END $$;`).Error; err != nil {
		return nil, fmt.Errorf("failed to create room_status enum: %v", err)
	}

	// Main connection: prefix all tables with room.
	db, err := gorm.Open(postgres.Open(cfg.DBURL), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "room.",
		},
	})
	if err != nil {
		return nil, fmt.Errorf("connect db failed: %v", err)
	}

	// search_path only room (no public)
	if err := db.Exec(`SET search_path TO room;`).Error; err != nil {
		return nil, fmt.Errorf("failed to set search_path: %v", err)
	}

	// Migrate (tables created as room.<table>)
	if err := db.AutoMigrate(
		&models.Room{},
		&models.Location{},
		&models.RoomImage{},
		&models.RoomType{},
		&models.Review{},
		&models.RoomBooking{},
		&models.Service{},
	); err != nil {
		return nil, fmt.Errorf("migration failed: %v", err)
	}

	return db, nil
}
