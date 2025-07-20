package database

import (
	"fmt"
	"my-app/internal/config"
	"my-app/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	if cfg.DBURL == "" {
		return nil, fmt.Errorf("DB_URL is not set")
	}

	db, err := gorm.Open(postgres.Open(cfg.DBURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("connect db failed : %v", err)
	}

	// Enable uuid extension
	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
		return nil, fmt.Errorf("enable uuid extension failed : %v", err)
	}

	// Enum Vip account
	if err := db.Exec(`DO $$ BEGIN
	CREATE TYPE vip_status AS ENUM ('nm','vp','2vp','3vp','b');
	EXCEPTION WHEN duplicate_object THEN NULL;
	END $$;`).Error; err != nil {
		return nil, fmt.Errorf("failed to create vip_status enum: %v", err)
	}
	// Auto migrate
	if err := db.AutoMigrate(&models.Role{}, &models.Account{}); err != nil {
		return nil, fmt.Errorf("migrate failed : %v", err)
	}
	return db, nil
}
