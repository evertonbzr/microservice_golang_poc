package main

import (
	"log/slog"
	"os"

	"github.com/evertonbzr/microservice_golang_poc/internal/api"
	"github.com/evertonbzr/microservice_golang_poc/internal/config"
	"github.com/evertonbzr/microservice_golang_poc/internal/db"
)

func main() {
	config.Load(os.Getenv("ENV"))

	slog.Info("Starting API...", "port", config.PORT, "env", config.ENV)

	db := db.Database(config.DATABASE_URL)
	slog.Info("Connected to database")

	apiCfg := &api.APIConfig{
		Port: config.PORT,
		DB:   db,
	}

	api.Start(apiCfg)
}
