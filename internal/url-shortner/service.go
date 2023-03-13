package url_shortner

import (
	"context"
	"github.com/amirhosseinmoayedi/URl-Shortener/internal/log"
)

type Service struct {
	URLRepo URLRepository
}

func NewService(repo URLRepository) Service {
	return Service{URLRepo: repo}
}

func (s Service) ShortenLink(ctx context.Context, url *URL) error {
	log.Logger.WithFields(map[string]interface{}{"ctx": ctx, "url": url}).Info("request for shorten the url")
	if err := url.validateOriginalPath(); err != nil {
		log.Logger.WithFields(map[string]interface{}{"ctx": ctx, "url": url}).Error(err)
		return err
	}
	if err := url.setCreateAt(); err != nil {
		log.Logger.WithFields(map[string]interface{}{"ctx": ctx, "url": url}).Error(err)
		return err
	}
	if err := url.setPath(); err != nil {
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

func (s Service) GetShortenLink(ctx context.Context, path string) (*URL, error) {
	log.Logger.WithFields(map[string]interface{}{"ctx": ctx, "path": path}).Info("request for get the original url")
	url, err := s.URLRepo.Find(ctx, path)
	if err != nil {
		log.Logger.WithFields(map[string]interface{}{"ctx": ctx, "path": path}).Info(err)
		return nil, err
	}
	log.Logger.WithFields(map[string]interface{}{"ctx": ctx, "path": path, "url": url}).Info("url short version found")
	return &url, nil
}
