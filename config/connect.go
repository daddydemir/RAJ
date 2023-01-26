package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/raj?parseTime=true"
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
