package repository

type Repository interface {
	SaveURL(original, short string) error
	GetOriginalURL(short string) (string, error)
	GetShortURL(original string) (string, error)
}
