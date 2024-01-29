package main

import (
	"github.com/LamichhaneBibek/go-template/internal/api"
	"github.com/LamichhaneBibek/go-template/internal/config"
	"github.com/LamichhaneBibek/go-template/internal/db"
	"github.com/LamichhaneBibek/go-template/internal/utils/logger"
)

func main() {
	config := config.GetConfig()
	log := logger.NewLogger(config)
	err := db.InitDB(config)
	defer db.CloseDB()
	if err != nil {
		log.Fatal(logger.Postgres, logger.Startup, err.Error(), nil)
	}
	db.Up()
	api.InitServer()
}
