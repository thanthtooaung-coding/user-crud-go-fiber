package api

import (
	"context"
	"github.com/thanthtooaung-coding/user-crud-go-fiber/handler"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func NewFiberApp(userHandler *handler.UserHandler) *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api")
	usersAPI := api.Group("/users")

	usersAPI.Post("/", userHandler.CreateUser)
	usersAPI.Get("/", userHandler.GetAllUsers)
	usersAPI.Get("/:id", userHandler.GetUserByID)

	return app
}

func RunServer(lc fx.Lifecycle, app *fiber.App) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Println("Fiber: Starting server on :3000")
				if err := app.Listen(":3000"); err != nil {
					log.Printf("Fiber: Failed to start server: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Fiber: Shutting down server...")
			return app.Shutdown()
		},
	})
}
