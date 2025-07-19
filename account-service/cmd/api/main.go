package main

import (
	"log"
	"my-app/internal/config"
	"my-app/internal/database"
	"my-app/internal/routes"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := database.InitDB(cfg)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	r := routes.SetupRouter(db)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
