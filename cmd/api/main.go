package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/evertonbzr/microservice_golang_poc/internal/api"
	"github.com/evertonbzr/microservice_golang_poc/internal/config"
	"github.com/evertonbzr/microservice_golang_poc/internal/db"
	"github.com/evertonbzr/microservice_golang_poc/internal/model"
)

func main() {
	config.Load(os.Getenv("ENV"))

	slog.Info("Starting API...", "port", config.PORT, "env", config.ENV)

	db := db.Database(config.DATABASE_URL)
	slog.Info("Connected to database")

	err := db.AutoMigrate(&model.Todo{})
	if err != nil {
		log.Fatal("Error migrating database", "error", err)
	}

	apiCfg := &api.APIConfig{
		Port: config.PORT,
		DB:   db,
	}

	api.Start(apiCfg)
}
