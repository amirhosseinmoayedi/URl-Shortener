package url_shortner

import (
	"context"
	"github.com/amirhosseinmoayedi/URl-Shortener/internal/log"
)

type Service struct {
	URLRepo URLRepository
}

func NewService(repo URLRepository) *Service {
	return &Service{URLRepo: repo}
}

func (s Service) ShortenLink(ctx context.Context, url *URL) error {
	if err := url.validateOriginalPath(); err != nil {
		log.Logger.WithFields(map[string]interface{}{"ctx": ctx, "url": url}).Info(err)
		return err
	}
	if err := url.setCreateAt(); err != nil {
		log.Logger.WithFields(map[string]interface{}{"ctx": ctx, "url": url}).Info(err)
		return err
	}
	if err := url.setPath(); err != nil {
		log.Logger.WithFields(map[string]interface{}{"ctx": ctx, "url": url}).Warn(err)
		return err
	}
	if err := s.URLRepo.Add(ctx, *url); err != nil {
		log.Logger.WithFields(map[string]interface{}{"ctx": ctx, "url": url}).Warn(err)
		return err
	}
	return nil
}

func (s Service) GetShortenLink(ctx context.Context, path string) (*URL, error) {
	url, err := s.URLRepo.Find(ctx, path)
	if err != nil {
		log.Logger.WithFields(map[string]interface{}{"ctx": ctx, "path": path}).Info(err)
		return nil, err
	}
	return &url, nil
}
