package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"url-shortener/internal/handler"
)

// Мокаем сервис для тестов
type MockService struct{}

func (m *MockService) Shorten(url string) (string, error) {
	// Возвращаем "сокращенный" URL для теста
	return "shortened-url", nil
}

func (m *MockService) Resolve(shortURL string) (string, error) {
	// Возвращаем оригинальный URL для теста
	return "http://example.com", nil
}

func TestShortenURL(t *testing.T) {
	// Инициализируем обработчик с мок-сервисом
	mockService := &MockService{}
	handler := handler.NewHandler(mockService)

	// Создаем тестовый HTTP-сервер с обработчиком
	recorder := httptest.NewRecorder()
	requestBody := map[string]string{"url": "http://example.com"}
	jsonData, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, "/shorten", bytes.NewReader(jsonData))

	// Вызываем обработчик
	handler.ShortenURL(recorder, request)

	// Проверяем, что код статуса 200 OK
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Проверяем, что ответ в JSON-формате и содержит ожидаемый "short_url"
	var response map[string]string
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "http://localhost:8080/shortened-url", response["short_url"])
}

func TestResolveURL(t *testing.T) {
	// Инициализируем обработчик с мок-сервисом
	mockService := &MockService{}
	handler := handler.NewHandler(mockService)

	// Создаем тестовый HTTP-сервер с обработчиком
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/shortened-url", nil)

	// Вызываем обработчик
	handler.ResolveURL(recorder, request)

	// Проверяем, что произошло перенаправление (HTTP код 302)
	assert.Equal(t, http.StatusFound, recorder.Code)

	// Проверяем, что заголовок Location содержит правильный URL
	location := recorder.Header().Get("Location")
	assert.Equal(t, "http://example.com", location)
}
