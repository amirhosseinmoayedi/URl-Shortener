package service

import (
	"context"
	"github.com/amirhosseinmoayedi/URl-Shortener/internal/domain/entity"
	"github.com/amirhosseinmoayedi/URl-Shortener/internal/domain/repository"
	"github.com/amirhosseinmoayedi/URl-Shortener/internal/log"
)

type Service struct {
	URLRepo repository.URLRepository
}

func NewService(repo repository.URLRepository) Service {
	return Service{URLRepo: repo}
}

func (s Service) Shortening(ctx context.Context, url *entity.URL) error {
	log.Logger.WithFields(map[string]interface{}{"ctx": ctx, "url": url}).Info("request for shorten the url")
	if err := url.ValidateOriginalPath(); err != nil {
		log.Logger.WithFields(map[string]interface{}{"ctx": ctx, "url": url}).Error(err)
		return err
	}
	if err := url.SetCreateAt(); err != nil {
		log.Logger.WithFields(map[string]interface{}{"ctx": ctx, "url": url}).Error(err)
		return err
	}
	if err := url.SetPath(); err != nil {
		log.Logger.WithFields(map[string]interface{}{"ctx": ctx, "url": url}).Error(err)
		return err
	}
	if err := s.URLRepo.Add(ctx, *url); err != nil {
		log.Logger.WithFields(map[string]interface{}{"ctx": ctx, "url": url}).Error(err)
		return err
	}
	log.Logger.WithFields(map[string]interface{}{"ctx": ctx, "url": url}).Info("url shorted")
	return nil
}

func (s Service) GetShortened(ctx context.Context, path string) (*entity.URL, error) {
	log.Logger.WithFields(map[string]interface{}{"ctx": ctx, "path": path}).Info("request for get the original url")
	url, err := s.URLRepo.Find(ctx, path)
	if err != nil {
		log.Logger.WithFields(map[string]interface{}{"ctx": ctx, "path": path}).Info(err)
		return nil, err
	}
	log.Logger.WithFields(map[string]interface{}{"ctx": ctx, "path": path, "url": url}).Info("url short version found")
	return &url, nil
}
