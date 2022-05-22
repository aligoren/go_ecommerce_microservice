package database

import (
	"fmt"
	"github.com/aligoren/go_ecommerce_microservice/auth-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func ConnectDb() {

	dbHost := os.Getenv("DB_HOST")
	userName := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	sslMode := os.Getenv("SSL_MODE")
	timeZone := os.Getenv("TIME_ZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", dbHost, userName, password, dbName, dbPort, sslMode, timeZone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Failed to connect to database %v\n, dsn: %s", err, dsn)
	}

	err = db.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatalf("Failed to auto migration %v", err)
	}

	DB = DbInstance{
		Db: db,
	}
}
