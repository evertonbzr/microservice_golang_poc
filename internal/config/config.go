package config

import (
	"log"
	"log/slog"

	"github.com/spf13/viper"
)

var (
	PORT         string
	ENV          string
	NAME         string
	DATABASE_URL string
)

func Load(env string) {
	slog.Info("Loading config...", "env", env)

	if env == "" {
		viper.Set("ENV", "development")
		viper.Set("PORT", "3000")
		viper.Set("NAME", "module_todo")
		viper.SetConfigFile(".env")

		if err := viper.ReadInConfig(); err != nil {
			log.Fatal("Error loading .env file", "error", err)
		}
	} else {
		viper.AutomaticEnv()
	}

	PORT = viper.GetString("PORT")
	ENV = viper.GetString("ENV")
	NAME = viper.GetString("NAME")
	DATABASE_URL = viper.GetString("DATABASE_URL")
}

func IsDevelopment() bool {
	return ENV == "development"
}

func IsProduction() bool {
	return ENV == "production"
}

func IsTest() bool {
	return ENV == "test"
}
