package main

import (
	"github.com/thanthtooaung-coding/user-crud-go-fiber/api"
	"github.com/thanthtooaung-coding/user-crud-go-fiber/handler"
	"github.com/thanthtooaung-coding/user-crud-go-fiber/internal/database"
	"go.uber.org/fx"
	"log"
)

func main() {
	app := fx.New(
		fx.Provide(
			database.NewGormDb,
			handler.NewUserHandler,
			api.NewFiberApp,
		),
		fx.Invoke(api.RunServer),
	)

	app.Run()

	if err := app.Err(); err != nil {
		log.Fatal(err)
	}
}
