package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/evertonbzr/microservice_golang_poc/internal/model"
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

	app.Route("/todos", func(r fiber.Router) {
		r.Get("/", func(c *fiber.Ctx) error {
			todos := []model.Todo{}

			if err := cfg.DB.Find(&todos).Error; err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": err.Error(),
				})
			}

			return c.JSON(fiber.Map{
				"todos": todos,
			})
		})

		r.Post("/", func(c *fiber.Ctx) error {
			todo := new(model.Todo)

			if err := c.BodyParser(todo); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err.Error(),
				})
			}

			if err := cfg.DB.Create(todo).Error; err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": err.Error(),
				})
			}

			return c.JSON(fiber.Map{
				"todo": todo,
			})
		})
	})

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
