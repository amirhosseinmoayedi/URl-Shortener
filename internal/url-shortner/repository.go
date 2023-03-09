package url_shortner

import (
	"context"
	"errors"
)

var URLNotFound = errors.New("URL not found")

type URLRepository interface {
	Add(ctx context.Context, url URL) error
	Find(ctx context.Context, path string) (URL, error)
}

type URLConvertor interface {
	ToURL() URL
}
