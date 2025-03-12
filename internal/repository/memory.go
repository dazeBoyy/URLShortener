package repository

import (
	"errors"
	"sync"
)

type InMemoryStorage struct {
	data map[string]string
	lock sync.RWMutex
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{data: make(map[string]string)}
}

func (s *InMemoryStorage) SaveURL(original, short string) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.data[short] = original
	return nil
}

func (s *InMemoryStorage) GetShortURL(original string) (string, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	for short, orig := range s.data {
		if orig == original {
			return short, nil
		}
	}
	return "", errors.New("URL not found")
}

func (s *InMemoryStorage) GetOriginalURL(short string) (string, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	original, exists := s.data[short]
	if !exists {
		return "", errors.New("URL not found")
	}
	return original, nil
}
