package cache

import (
	"context"
	"errors"
	"time"
	"url-shortener/internal/url-shortner"
)

type InMemoryCacheURLRepository struct {
	cache map[string]InMemoryCacheURL
}

type InMemoryCacheURL struct {
	Path      string
	Original  string
	CreatedAt time.Time
}

func NewInMemoryCacheURLRepository() *InMemoryCacheURLRepository {
	cache := make(map[string]InMemoryCacheURL)
	return &InMemoryCacheURLRepository{cache: cache}
}

func (ir *InMemoryCacheURLRepository) Add(ctx context.Context, url url_shortner.URL) error {
	imURL := ToInMemoryCacheURL(url)
	if _, ok := ir.cache[url.Path]; ok {
		return errors.New("URL already exists")
	}
	ir.cache[url.Path] = imURL
	return nil
}

func (ir *InMemoryCacheURLRepository) Find(ctx context.Context, path string) (url_shortner.URL, error) {
	v, ok := ir.cache[path]
	if !ok {
		return url_shortner.URL{}, url_shortner.URLNotFound
	}
	url := v.ToURL()
	return url, nil
}

// ToURL is a DTO
func (i *InMemoryCacheURL) ToURL() url_shortner.URL {
	return url_shortner.URL{
		Path:      i.Path,
		Original:  i.Original,
		CreatedAt: i.CreatedAt,
	}
}

func ToInMemoryCacheURL(url url_shortner.URL) InMemoryCacheURL {
	return InMemoryCacheURL{
		Path:      url.Path,
		Original:  url.Original,
		CreatedAt: url.CreatedAt,
	}
}
