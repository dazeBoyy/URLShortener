package service

import (
	_ "errors"
	"math/rand"
	"url-shortener/internal/repository"
)

const shortURLLength = 10
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

type URLService struct {
	repo repository.Repository
}

func NewURLService(repo repository.Repository) *URLService {
	return &URLService{repo: repo}
}

func (s *URLService) Shorten(original string) (string, error) {
	if existing, err := s.repo.GetShortURL(original); err == nil {
		return existing, nil
	}

	short := generateShortURL()
	if err := s.repo.SaveURL(original, short); err != nil {
		return "", err
	}
	return short, nil
}

func (s *URLService) Resolve(short string) (string, error) {
	return s.repo.GetOriginalURL(short)
}

func generateShortURL() string {
	result := make([]byte, shortURLLength)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
