package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	client := http.Client{}
	ctx := context.Background()

	url, err := CreateUrl(os.Args[1], apiKey)
	if err != nil {
		log.Fatal(err)
	}

	dbConfig, err := NewDBConfig()
	if err != nil {
		log.Fatal(err)
	}

	rdb, err := DBInit(ctx, dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	var data APIResponse

	dbCtx, cancelDB := context.WithTimeout(ctx, time.Second)
	defer cancelDB()

	dataStr, err := rdb.Get(dbCtx, url).Result()
	if err != redis.Nil {
		err = json.Unmarshal([]byte(dataStr), &data)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		resp, err := client.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		err = rdb.Set(dbCtx, url, string(body), 12*time.Hour).Err()
		if err != nil {
			fmt.Printf("%s\n", err.Error())
		}

		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Fatal(err)
		}
	}

	displayTab(data)

}
