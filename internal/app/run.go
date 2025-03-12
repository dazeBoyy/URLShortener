package app

import (
	"fmt"
	"net/http"
	"os"

	"url-shortener/internal/handler"
	"url-shortener/internal/repository"
	"url-shortener/internal/service"
)

func Run() error {
	var repo repository.Repository
	// Получаем значение переменной окружения STORAGE для выбора типа хранилища.
	storageType := os.Getenv("STORAGE")

	if storageType == "postgres" {
		// Если STORAGE = postgres, используем PostgreSQL
		dbURL := os.Getenv("DATABASE_URL")
		if dbURL == "" {
			return fmt.Errorf("DATABASE_URL is required when using PostgreSQL")
		}
		fmt.Println("Using PostgreSQL storage")
		ps, err := repository.NewPostgresStorage(dbURL)
		if err != nil {
			return fmt.Errorf("failed to connect to PostgreSQL: %w", err)
		}
		repo = ps
	} else if storageType == "memory" {
		// Если STORAGE != postgres или не задано, используем in-memory storage
		fmt.Println("Using in-memory storage")
		repo = repository.NewInMemoryStorage()
	} else {
		// Если STORAGE не задан или задан некорректно, возвращаем ошибку
		return fmt.Errorf("invalid storage type: %s, use 'postgres' or 'memory'", storageType)
	}

	// Создаем экземпляр сервиса для сокращения ссылок, но присваиваем переменной с типом интерфейса для гибкости.
	var shortener service.URLShortenerInter = service.NewURLService(repo)
	handler := handler.NewHandler(shortener)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	fmt.Println("Server is running on port 8080...")
	return http.ListenAndServe(":8080", mux)
}
