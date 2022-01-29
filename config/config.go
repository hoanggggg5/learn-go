package config

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Store = session.New()
var Database *gorm.DB

func InitConfig() {
	dsn := "host=localhost user=postgres password=hoang4869 dbname=learn-fiber port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	Database = db
}
