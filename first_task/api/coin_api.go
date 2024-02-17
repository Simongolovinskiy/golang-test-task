package api

import (
	"context"
	"github.com/carlmjohnson/requests"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type CoinApi struct{}

func getBaseUrl() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file couldn't be loaded")
	}
	baseUrl := os.Getenv("FIRST_TASK_BASE_URL")
	return baseUrl
}

func (c CoinApi) GetCoins() string {
	var response string
	baseUrl := getBaseUrl()
	err := requests.
		URL(baseUrl).
		ToString(&response).
		Fetch(context.Background())
	if err != nil {
		log.Fatal("could not connect to %s", baseUrl)
	}
	return response
}
