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
	"user-crud-service/internal/database"
	"user-crud-service/internal/handler"
	"user-crud-service/repository"
	"user-crud-service/server"
	"user-crud-service/service"
)

func main() {
	// todo write README, about launch
	// todo docs

	// todo tests

	// todo login
	// todo double login & session
	// todo auth protection

	// todo POST concurrency
	// todo mongo transactions
	// todo handle if user by ID already exists (addUser) | unique entries in db

	// todo loglevel
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

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

	//userRepo := repository.NewUserRepository(conn.Database("users"))
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
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

	defer func() {
		cancel()
		defer db.Close(ctx)
		//defer serv.Stop()
	}()

	// Shutdown
	err = serv.Shutdown(ctx)
	if err != nil {
		log.Error().Msgf("server shutdown err: %v", err)
	}
}
