package url_shortner

import (
	"context"
	"errors"
)

type Service struct {
	URLRepo URLRepository
}

func NewService(repo URLRepository) (*Service, error) {
	if repo == nil {
		return nil, errors.New("URL repository cant be nil")
	}
	return &Service{URLRepo: repo}, nil
}

func (s Service) ShortenLink(ctx context.Context, url *URL) error {
	if err := url.validateOriginalPath(); err != nil {
		return err
	}
	if err := url.setCreateAt(); err != nil {
		return err
	}
	if err := url.setPath(); err != nil {
		return err
	}
	if err := s.URLRepo.Add(ctx, *url); err != nil {
		return err
	}
	return nil
}

func (s Service) GetShortenLink(ctx context.Context, path string) (*URL, error) {
	url, err := s.URLRepo.Find(ctx, path)
	if err != nil {
		return nil, err
	}
	return &url, nil
}
