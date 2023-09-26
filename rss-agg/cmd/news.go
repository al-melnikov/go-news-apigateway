package main

import (
	"encoding/json"
	"log"
	"os"
	"rss-agg/pkg/models"
	db "rss-agg/pkg/postgres"
	"rss-agg/pkg/rss"
	"time"

	"github.com/joho/godotenv"
)

// конфигурация приложения
type config struct {
	URLS   []string `json:"rss"`
	Period int      `json:"request_period"`
}

func main() {
	log.Print("server has started")

	godotenv.Load()

	//get the port from the environment variable
	dbString := os.Getenv("DB_URL")
	if dbString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	// чтение и раскодирование файла конфигурации
	b, err := os.ReadFile("./config.json")

	if err != nil {
		log.Fatal(err)
	}
	var config config
	err = json.Unmarshal(b, &config)
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.New(dbString)
	if err != nil {
		log.Fatal(err)
	}

	chPosts := make(chan []models.Post)
	chErrs := make(chan error)
	for _, url := range config.URLS {
		go parseURL(url, chPosts, chErrs, config.Period)
	}

	// запись потока новостей в БД
	go func() {
		for posts := range chPosts {
			for _, post := range posts {
				db.AddPost(post)
			}
		}
	}()

	// обработка потока ошибок

	go func() {
		for err := range chErrs {
			log.Println("ошибка:", err)
		}
	}()

	// запись потока новостей в БД

	for posts := range chPosts {
		for _, post := range posts {
			db.AddPost(post)
		}
	}

}

func parseURL(url string, posts chan<- []models.Post, errs chan<- error, period int) {
	for {
		news, err := rss.RssToStruct(url)
		if err != nil {
			errs <- err
			continue
		}
		posts <- news
		time.Sleep(time.Minute * time.Duration(period))
	}
}
