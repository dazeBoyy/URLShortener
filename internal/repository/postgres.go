package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	DB *sql.DB
}

func NewPostgresStorage(dsn string) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	query := `
	CREATE TABLE IF NOT EXISTS urls (
		id SERIAL PRIMARY KEY,
		original_url TEXT UNIQUE NOT NULL,
		short_url VARCHAR(10) UNIQUE NOT NULL
	);
	`
	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	return &PostgresStorage{DB: db}, nil
}

func (s *PostgresStorage) SaveURL(original, short string) error {
	_, err := s.DB.Exec("INSERT INTO urls (original_url, short_url) VALUES ($1, $2) ON CONFLICT (original_url) DO NOTHING", original, short)
	return err
}

func (s *PostgresStorage) GetShortURL(original string) (string, error) {
	var short string
	err := s.DB.QueryRow("SELECT short_url FROM urls WHERE original_url = $1", original).Scan(&short)
	if err != nil {
		return "", err
	}
	return short, nil
}

func (s *PostgresStorage) GetOriginalURL(short string) (string, error) {
	var original string
	err := s.DB.QueryRow("SELECT original_url FROM urls WHERE short_url = $1", short).Scan(&original)
	if err != nil {
		return "", err
	}
	return original, nil
}
