package memory

import (
	"context"
	"errors"
	"github.com/amirhosseinmoayedi/URl-Shortener/internal/domain/entity"
	"github.com/amirhosseinmoayedi/URl-Shortener/internal/domain/repository"
	"time"
)

type InMemoryURLRepository struct {
	cache map[string]InMemoryURL
}

type InMemoryURL struct {
	Path      string
	Original  string
	CreatedAt time.Time
}

func NewInMemoryCacheURLRepository() *InMemoryURLRepository {
	cache := make(map[string]InMemoryURL)
	return &InMemoryURLRepository{cache: cache}
}

func (ir *InMemoryURLRepository) Add(ctx context.Context, url entity.URL) error {
	imURL := InMemoryURL{
		Path:      url.Path,
		Original:  url.Original,
		CreatedAt: url.CreatedAt,
	}
	if _, ok := ir.cache[url.Path]; ok {
		return errors.New("URL already exists")
	}
	ir.cache[url.Path] = imURL
	return nil
}

func (ir *InMemoryURLRepository) Find(ctx context.Context, path string) (entity.URL, error) {
	imURL, ok := ir.cache[path]
	if !ok {
		return entity.URL{}, repository.URLNotFound
	}
	url := entity.URL{
		Path:      imURL.Path,
		Original:  imURL.Original,
		CreatedAt: imURL.CreatedAt,
	}
	return url, nil
}
