package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	dsn := Get("DB_DSN")
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
