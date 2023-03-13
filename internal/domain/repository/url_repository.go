package repository

import (
	"context"
	"errors"
	"github.com/amirhosseinmoayedi/URl-Shortener/internal/domain/entity"
)

var URLNotFound = errors.New("URL not found")
var URLAlreadyExists = errors.New("URL already exists")

type URLRepository interface {
	Add(ctx context.Context, url entity.URL) error
	Find(ctx context.Context, path string) (entity.URL, error)
}
