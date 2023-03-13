package postgres

import (
	"context"
	"errors"
	"github.com/amirhosseinmoayedi/URl-Shortener/internal/domain/entity"
	"github.com/amirhosseinmoayedi/URl-Shortener/internal/domain/repository"
	"gorm.io/gorm"
)

type URLRepositoryPostgres struct {
	db *gorm.DB
}

type URLPostgres struct {
	gorm.Model
	Path     string `gorm:"uniqueIndex:idx_path,not null"`
	Original string `gorm:"unique"`
}

func NewPostgresURLRepository(db *gorm.DB) *URLRepositoryPostgres {
	return &URLRepositoryPostgres{
		db: db,
	}
}

func (p *URLRepositoryPostgres) Add(ctx context.Context, url entity.URL) error {
	postgresURl := URLPostgres{
		Path:     url.Path,
		Original: url.Original,
	}

	db := p.db.WithContext(ctx)

	result := db.Create(&postgresURl)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return repository.URLAlreadyExists
		}
		return result.Error
	}
	return nil
}

func (p *URLRepositoryPostgres) Find(ctx context.Context, path string) (entity.URL, error) {
	var url URLPostgres
	db := p.db.WithContext(ctx)
	result := db.Where("path = ?", path).First(&url)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entity.URL{}, repository.URLNotFound
		}
		return entity.URL{}, result.Error
	}

	return entity.URL{Path: url.Path, Original: url.Original}, nil
}
