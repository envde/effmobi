package main

// @title           Effmobi Subscriptions API
// @version         1.0
// @description     API для управления подписками пользователей
// @host            localhost:8080
// @BasePath        /

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/envde/effmobi/interanl/app"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file, reading from environment")
	}

	app.Run()
}
