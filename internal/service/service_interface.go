package service

type URLShortenerInter interface {
	Shorten(original string) (string, error)
	Resolve(short string) (string, error)
}
