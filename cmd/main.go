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
    "user-crud-service/internal/repository"
    "user-crud-service/internal/server"
    "user-crud-service/internal/service"
)

func main() {
    // todo add app pkg: app.Run()
    // todo mv config to cmd, add default configs
    // todo add panic recovery etc
    // todo tests

    zerolog.SetGlobalLevel(zerolog.ErrorLevel)

    cfg, err := config.NewConfig("config/.config.env")
    if err != nil {
        log.Error().Msg("Initial configuration failed")
        cfg, err = config.NewConfig("default")
        if err != nil {
            log.Fatal().Msg("Default configuration failed")
        }
    }

    db, err := database.New()
    if err != nil {
        log.Error().Msg(fmt.Sprintf("failed db conn: %v", err))
        return
    }

    userRepo := repository.NewUserRepository(db)
    userService := service.NewUserService(userRepo)
    //userHandler := handler.NewUserHandler(userService)
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
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer func() {
        cancel()
        //defer db.Close()
        //    defer server.Stop()
    }()

    // Shutdown
    err = serv.Shutdown(ctx)
    if err != nil {
        log.Error().Msg(fmt.Sprintf("server shutdown err: %v", err))
    }
}
