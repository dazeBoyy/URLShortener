package main

import (
	"github.com/joho/godotenv"
	"log"
	"url-shortener/internal/app"
)

func main() {
	// Загружаем переменные из .env файла
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Запускаем приложение
	if err := app.Run(); err != nil {
		log.Fatalf("Error starting application: %v", err)
	}
}
