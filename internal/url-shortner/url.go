package url_shortner

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"regexp"
	"time"
)

var domainRegex *regexp.Regexp

func init() {
	var err error
	domainRegex, err = regexp.Compile("^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9]\\.[a-zA-Z]{2,}$")
	if err != nil {
		err = fmt.Errorf("can't start the url shortner module because :%w", err)
		log.Fatal(err)
	}

}

// URL is an entity
type URL struct {
	// Path is the ID for url
	Path      string
	Original  string
	CreatedAt time.Time
}

func (u *URL) validateOriginalPath() error {
	if !domainRegex.MatchString(u.Original) {
		return errors.New("URL is not valid")
	}
	return nil
}

func (u *URL) setCreateAt() error {
	if u.CreatedAt.IsZero() {
		return errors.New("URL already have created at")
	}
	u.CreatedAt = time.Now()
	return nil
}

func (u *URL) setPath() error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	u.Path = id.String()
	return nil
}

type Service struct {
	URLRepo URLRepository
}

func NewService(repo URLRepository) *Service {
	return &Service{URLRepo: repo}
}

func (s Service) ShortenLink(ctx context.Context, url *URL) error {
	if err := url.validateOriginalPath(); err != nil {
		return err
	}
	if err := url.setCreateAt(); err != nil {
		return err
	}
	if err := s.URLRepo.add(ctx, *url); err != nil {
		return err
	}
	return nil
}

func (s Service) GetShortenLink(ctx context.Context, path string) (*URL, error) {
	url, err := s.URLRepo.find(ctx, path)
	if err != nil {
		return nil, err
	}
	return url, nil
}

//feat: add domain service for urlshoten
//
//URL entity and validation methods addeed
//domain service for setting url
