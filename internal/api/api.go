package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/valyala/fasthttp"

	"gorm.io/gorm"
)

type APIConfig struct {
	DB   *gorm.DB
	Port string
}

func Start(cfg *APIConfig) {
	app := fiber.New()
	app.Use(healthcheck.New())
	app.Use(requestid.New())
	app.Use(logger.New())
	app.Use(recover.New())

	srv := fasthttp.Server{
		Handler: app.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(":" + cfg.Port); err != nil && err != http.ErrServerClosed {
			log.Fatal("Error starting server: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.ShutdownWithContext(ctx); err != nil {
		log.Fatal("Error shutting down server: ", err)
	}

	select {
	case <-ctx.Done():
		log.Println("Timeout of 5 seconds.")
	default:
		log.Println("Server gracefully stopped.")
	}
}
