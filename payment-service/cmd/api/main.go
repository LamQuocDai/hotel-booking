package main

import (
	"log"
	"payment-service/internal/config"
	"payment-service/internal/database"
	"payment-service/internal/routes"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
		return
	}

	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect db: %v", err)
		return
	}
	r := routes.SetupRouter(db)
	if err := r.Run(":8083"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
		return
	}
}
