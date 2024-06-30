package config

import (
	"github.com/ikotun/chrgo/internals/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dbUrl = "host=localhost user=postgres password=postgres dbname=chrgo port=5432 sslmode=disable"

var DB *gorm.DB

func InitDB() {
	DB, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	log.Info("Connected to database successfully")
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Session{})
	DB.AutoMigrate(&models.Message{})
	log.Info("Migrated models successfully")
	if err != nil {
		panic(err)
	}
	// DB = db

}
