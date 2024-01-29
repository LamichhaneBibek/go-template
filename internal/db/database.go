package db

import (
	"fmt"
	"log"

	"github.com/LamichhaneBibek/go-template/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func InitDB(config *config.Config) error {
	var err error
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Kathmandu",
		config.Postgres.Host, config.Postgres.Port, config.Postgres.User, config.Postgres.Password,
		config.Postgres.DatabaseName, config.Postgres.SSLMode)
	dbClient, err = gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDb, _ := dbClient.DB()
	err = sqlDb.Ping()
	if err != nil {
		return err
	}

	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)

	log.Println("Database connection established")
	return nil
}

func GetDB() *gorm.DB {
	return dbClient
}

func CloseDB() {
	sqlDb, _ := dbClient.DB()
	sqlDb.Close()
}
