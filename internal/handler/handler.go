package handler

import (
	"encoding/json"
	"net/http"
	"url-shortener/internal/service"
)

type Handler struct {
	service service.URLShortenerInter
}

func NewHandler(service service.URLShortenerInter) *Handler {
	return &Handler{service: service}
}

func (h *Handler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	shortURL, err := h.service.Shorten(request.URL)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"short_url": "http://localhost:8080/" + shortURL}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) ResolveURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	shortURL := r.URL.Path[1:] // Извлекаем сокращенный код из URL
	originalURL, err := h.service.Resolve(shortURL)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /shorten", h.ShortenURL) // POST-запрос для сокращения ссылки
	mux.HandleFunc("GET /", h.ResolveURL)         // GET-запрос для редиректа
}
