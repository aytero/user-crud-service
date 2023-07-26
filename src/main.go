package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"user-crud-service/config"
	"user-crud-service/database"
	"user-crud-service/handler"
	"user-crud-service/repository"
	"user-crud-service/server"
	"user-crud-service/service"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	cfg, err := config.NewConfig("config/.config.env")
	if err != nil {
		log.Error().Msg("initial configuration failed. Starting default configuration")
		cfg, err = config.NewConfig("default")
		if err != nil {
			log.Fatal().Msg("default configuration failed")
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	db, err := database.New(ctx, cfg.MG)
	if err != nil {
		log.Fatal().Msgf("failed db conn: %v", err)
	}

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(cfg.SRVC, userRepo)
	router := gin.New()
	handler.SetupUserHandler(router, userService)

	serv := server.NewServer(cfg.HTTP, router)
	serv.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	select {
	case sig := <-interrupt:
		log.Info().Msgf("shutting down with signal: %s\n", sig)
	case err = <-serv.Notify():
		log.Error().Msgf("server err: %s", err)
	}

	// Shutdown
	defer cancel()
	defer db.Close(ctx)
	err = serv.Shutdown(ctx)
	if err != nil {
		log.Error().Msgf("server shutdown err: %v", err)
	}
}
