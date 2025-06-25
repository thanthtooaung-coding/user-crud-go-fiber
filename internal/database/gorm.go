package database

import (
	"context"
	"fmt"
	"github.com/thanthtooaung-coding/user-crud-go-fiber/internal/domain/user"
	"go.uber.org/fx"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func NewGormDb(lc fx.Lifecycle) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("Migrating")
			return db.AutoMigrate(&user.User{})
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Closing")
			sqlDB, err := db.DB()
			if err != nil {
				return fmt.Errorf("failed to get DB: %w", err)
			}
			return sqlDB.Close()
		},
	})
	return db, nil
}
