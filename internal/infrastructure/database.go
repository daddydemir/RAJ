package infrastructure

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
)

var DB *gorm.DB

func InitDB(dsn string) {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("database", "message", err.Error())
		panic(err)
	}
	slog.Info("database", "message", "connected")
}

func GetDB() *gorm.DB {
	return DB
}
