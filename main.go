package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

type days map[string]any

func main() {
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	args := os.Args
	fmt.Println(args)
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s/?unitGroup=metric&key=%s", args[1], apiKey)
	client := http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data days
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data["days"])
}
