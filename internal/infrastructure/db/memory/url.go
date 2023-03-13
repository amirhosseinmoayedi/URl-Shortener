package memory

import (
	"context"
	"github.com/amirhosseinmoayedi/URl-Shortener/internal/domain/entity"
	"github.com/amirhosseinmoayedi/URl-Shortener/internal/domain/repository"
	"time"
)

type URLRepositoryInMemory struct {
	cache map[string]URLInMemory
}

type URLInMemory struct {
	Path      string
	Original  string
	CreatedAt time.Time
}

func NewURLRepositoryInMemory() *URLRepositoryInMemory {
	cache := make(map[string]URLInMemory)
	return &URLRepositoryInMemory{cache: cache}
}

func (ir *URLRepositoryInMemory) Add(ctx context.Context, url entity.URL) error {
	imURL := URLInMemory{
		Path:      url.Path,
		Original:  url.Original,
		CreatedAt: url.CreatedAt,
	}
	if _, ok := ir.cache[url.Path]; ok {
		return repository.URLAlreadyExists
	}
	ir.cache[url.Path] = imURL
	return nil
}

func (ir *URLRepositoryInMemory) Find(ctx context.Context, path string) (entity.URL, error) {
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
