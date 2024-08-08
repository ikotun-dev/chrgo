package config

import (
	"github.com/ikotun/chrgo/internals/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dbUrl = "host=localhost user=postgres password=postgres dbname=chr port=5432 sslmode=disable"

var DB *gorm.DB

func InitDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	log.Info("Connected to database successfully")
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Session{})
	db.AutoMigrate(&models.Message{})
	log.Info("Migrated models successfully")
	if err != nil {
		panic(err)
	}
	DB = db
	return DB

}
