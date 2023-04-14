package db

import (
	"codezas-pos/entity"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Connection *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)},
	)
	if err != nil {
		log.Fatal("Cannot connect to the database")
	}

	Connection = db
}

func Migrate() {
	Connection.AutoMigrate(
		&entity.Category{},
		&entity.Product{},
		&entity.Order{},
		&entity.OrderItem{},
	)
}
