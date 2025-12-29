package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"os"
	"strconv"
)

type Config struct {
	Addr     string
	Password string
	DB       int
}

func NewDBConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("unabled to create config: %w", err)
	}

	addr := os.Getenv("DB_ADDR")
	pwd := os.Getenv("DB_PASSWORD")
	dbType, err := strconv.Atoi(os.Getenv("DB_TYPE"))
	if err != nil {
		return nil, fmt.Errorf("unable to fetch env variable: %w", err)
	}

	return &Config{
		Addr:     addr,
		Password: pwd,
		DB:       dbType,
	}, nil
}

func DBInit(ctx context.Context, c *Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.DB,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
