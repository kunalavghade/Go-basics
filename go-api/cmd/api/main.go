package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kunalavghade/Go-basics/go-api/internal/config"
	"github.com/kunalavghade/Go-basics/go-api/internal/database"
	"github.com/kunalavghade/Go-basics/go-api/internal/logger"
)

func main() {
	fmt.Println("Hello World")
	log := logger.New("debug")
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}
	db, err := database.New(cfg.Database)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}
	mainDb, err := db.DB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get database connection")
	}
	err = mainDb.Ping()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to ping database")
	}
	defer mainDb.Close()
	log.Info().Msg("Connected to database")

	gin.SetMode(cfg.Server.GinMode)
	log.Info().Msg("Starting the server...")

}
