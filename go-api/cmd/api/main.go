package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kunalavghade/Go-basics/go-api/internal/config"
	"github.com/kunalavghade/Go-basics/go-api/internal/database"
	"github.com/kunalavghade/Go-basics/go-api/internal/logger"
	"github.com/kunalavghade/Go-basics/go-api/internal/server"
)

func main() {
	fmt.Println("Hello World")
	log := logger.New("debug")
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}
	db, err := database.New(&cfg.Database)
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

	srv := server.NewServer(cfg, &log, db)
	handler := srv.SetupRoutes()

	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Server.Port),
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Info().Str("port", httpServer.Addr).Msg("Starting the server...")
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("Failed to start the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutting down server gracefully...")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server forced to shutdown")
	}
	log.Info().Msg("Server exited")
}
