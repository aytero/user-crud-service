package config

import (
    "fmt"
    "github.com/joho/godotenv"
    "os"
)

type Config struct {
    HTTP
    Log
}

type (
    HTTP struct {
        Port     string
        LogLevel string
    }

    MG struct {
        // todo add Mongo config
    }

    Log struct {
        Level string
    }
)

func NewConfig(config string) (*Config, error) {
    cfg := &Config{}

    if config == "default" {
        cfg.Log.Level = "ERROR"
        cfg.HTTP.Port = "8080"
        cfg.HTTP.LogLevel = "ERROR"
        return cfg, nil
    }

    err := godotenv.Load(config)
    if err != nil {
        return nil, fmt.Errorf("config error: %w", err)
    }
    cfg.Log.Level = os.Getenv("LOG_LEVEL")
    cfg.HTTP.Port = os.Getenv("PORT")
    cfg.HTTP.LogLevel = cfg.Log.Level

    return cfg, nil
}
