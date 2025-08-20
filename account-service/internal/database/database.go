package database

import (
	"fmt"
	"log"
	"my-app/internal/config"
	"my-app/internal/models"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	if cfg.DBURL == "" {
		return nil, fmt.Errorf("DB_URL is not set")
	}

	// db, err := gorm.Open(postgres.Open(cfg.DBURL), &gorm.Config{})
	// if err != nil {
	// 	return nil, fmt.Errorf("connect db failed : %v", err)
	// }

	// configure GORM logger to reduce “SLOW SQL” noise
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             1500 * time.Millisecond, // increase from 200ms
			LogLevel:                  logger.Warn,             // Warn/Error to reduce chatter
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true, // less log spam
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.Open(cfg.DBURL), &gorm.Config{
		Logger:      newLogger,
		PrepareStmt: true, // speed up repeated queries
	})
	if err != nil {
		return nil, fmt.Errorf("connect db failed : %v", err)
	}

	// tune connection pool (PgBouncer/Supabase friendly)
	sqlDB, err := db.DB()
	if err == nil {
		sqlDB.SetMaxIdleConns(5)
		sqlDB.SetMaxOpenConns(10)
		sqlDB.SetConnMaxLifetime(30 * time.Minute)
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
	if err := db.AutoMigrate(&models.Role{}, &models.Account{}, &models.Permission{}, &models.DetailRole{}); err != nil {
		return nil, fmt.Errorf("migrate failed : %v", err)
	}
	return db, nil
}
