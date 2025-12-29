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
)

func main() {
	ctx := context.Background()
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	client := http.Client{}

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
	dataStr, err := rdb.Get(ctx, url).Result()
	if err != redis.Nil {
		fmt.Println("cache hit!")
		err = json.Unmarshal([]byte(dataStr), &data)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("fetching data from api...")
		resp, err := client.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		err = rdb.Set(ctx, url, string(body), 0).Err()
		if err != nil {
			fmt.Printf("%s\n", err.Error())
		}

		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, d := range data.Days {
		fmt.Printf("%s - %.f\n", d.DateTime, d.Temp)
	}
}
