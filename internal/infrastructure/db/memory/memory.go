package memory

import (
	"context"
	"errors"
	"github.com/amirhosseinmoayedi/URl-Shortener/internal/domain/entity"
	"github.com/amirhosseinmoayedi/URl-Shortener/internal/domain/repository"
	"time"
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

func (ir *InMemoryCacheURLRepository) Add(ctx context.Context, url entity.URL) error {
	imURL := ToInMemoryCacheURL(url)
	if _, ok := ir.cache[url.Path]; ok {
		return errors.New("URL already exists")
	}
	ir.cache[url.Path] = imURL
	return nil
}

func (ir *InMemoryCacheURLRepository) Find(ctx context.Context, path string) (entity.URL, error) {
	v, ok := ir.cache[path]
	if !ok {
		return entity.URL{}, repository.URLNotFound
	}
	url := v.ToURL()
	return url, nil
}

// ToURL is a DTO
func (i *InMemoryCacheURL) ToURL() entity.URL {
	return entity.URL{
		Path:      i.Path,
		Original:  i.Original,
		CreatedAt: i.CreatedAt,
	}
}

func ToInMemoryCacheURL(url entity.URL) InMemoryCacheURL {
	return InMemoryCacheURL{
		Path:      url.Path,
		Original:  url.Original,
		CreatedAt: url.CreatedAt,
	}
}
