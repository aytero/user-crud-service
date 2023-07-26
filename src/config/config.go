package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	HTTP
	MG
	Log
	SRVC
}

type (
	HTTP struct {
		Port     string
		LogLevel string
	}

	MG struct {
		User           string
		Password       string
		Host           string
		DbName         string
		CollectionName string
	}

	Log struct {
		Level string
	}

	SRVC struct {
		FileLocation string
	}
)

func NewConfig(config string) (*Config, error) {
	cfg := &Config{}

	if config == "default" {
		cfg.Log.Level = "ERROR"
		cfg.HTTP.Port = "8080"
		cfg.HTTP.LogLevel = "ERROR"
		cfg.MG.User = "root"
		cfg.MG.Password = "password"
		cfg.MG.Host = "localhost"
		cfg.MG.DbName = "users"
		cfg.MG.CollectionName = "users"
		return cfg, nil
	}

	err := godotenv.Load(config)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}
	cfg.Log.Level = os.Getenv("LOG_LEVEL")
	cfg.HTTP.Port = os.Getenv("PORT")
	cfg.HTTP.LogLevel = cfg.Log.Level

	cfg.MG.User = os.Getenv("MONGO_USER")
	cfg.MG.Password = os.Getenv("MONGO_PASSWORD")
	cfg.MG.Host = os.Getenv("MONGO_HOST")
	cfg.MG.DbName = os.Getenv("MONGO_DBNAME")
	cfg.MG.CollectionName = os.Getenv("MONGO_COLLECTION_NAME")

	cfg.SRVC.FileLocation = os.Getenv("FILE_LOCATION")

	return cfg, nil
}
