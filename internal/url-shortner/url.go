package url_shortner

import (
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

func NewUrl(originalURL string) *URL {
	return &URL{Original: originalURL}
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
