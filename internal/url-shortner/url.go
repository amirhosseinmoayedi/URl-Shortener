package url_shortner

import (
	"errors"
	"fmt"
	log "github.com/amirhosseinmoayedi/URl-Shortener/internal/log"
	"hash/fnv"
	"regexp"
	"strconv"
	"time"
)

var domainRegex *regexp.Regexp

func init() {
	var err error
	domainRegex, err = regexp.Compile("^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9]\\.[a-zA-Z]{2,}$")
	if err != nil {
		err = fmt.Errorf("can't start the url shortner module because :%w", err)
		log.Logger.Fatal(err)
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
	if !u.CreatedAt.IsZero() {
		return errors.New("URL already have created at")
	}
	u.CreatedAt = time.Now()
	return nil
}

func (u *URL) setPath() error {
	h := fnv.New32()
	_, err := h.Write([]byte(u.Original))
	if err != nil {
		log.Logger.WithFields(map[string]interface{}{"hash": h, "url": u.Original}).Warn(err)
		return err
	}
	u.Path = strconv.Itoa(int(h.Sum32()))
	return nil
}
