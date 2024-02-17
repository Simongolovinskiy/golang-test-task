package api

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type ApiParser struct{}

func getBaseUrl() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file couldn't be loaded")
	}
	baseUrl := os.Getenv("SECOND_TASK_BASE_URL")
	return baseUrl
}

func (c ApiParser) GetData() *goquery.Document {

	baseUrl := getBaseUrl()
	response, err := http.Get(baseUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}
