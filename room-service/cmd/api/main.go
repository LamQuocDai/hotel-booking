package main

import (
	"log"
	"room-service/internal/config"
	"room-service/internal/database"
	"room-service/internal/routes"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	db, err := database.InitDB(cfg)

	if err != nil {
		log.Fatalf("Failed to connect to db: %v", err)
	}

	r := routes.SetupRouter(db)
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
