package main

import (
	"context"
	"fmt"
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
	// todo add app pkg: app.Run()
	// todo mv config to cmd, add default configs
	// todo add panic recovery etc
	// todo better logger

	// todo tests
	// todo password hashing
	// todo working with documents: upload/update/delete
	// todo write README
	// todo login
	// todo read about PATCH DELETE UPDATE requests
	// todo handle if user by ID already exists (addUser) | unique entries in db
	// todo docs

	// todo loglevel
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	cfg, err := config.NewConfig("config/.config.env")
	if err != nil {
		log.Error().Msg("Initial configuration failed. Starting default configuration")
		cfg, err = config.NewConfig("default")
		if err != nil {
			log.Fatal().Msg("Default configuration failed")
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	db, err := database.New(ctx, cfg.MG)
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("failed db conn: %v", err))
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
		log.Printf("shutting down with signal: %s\n", sig)
	case err = <-serv.Notify():
		log.Printf("server err: %s", err)
	}

	defer func() {
		cancel()
		defer db.Close(ctx)
		//defer serv.Stop()
	}()

	// Shutdown
	err = serv.Shutdown(ctx)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("server shutdown err: %v", err))
	}
}