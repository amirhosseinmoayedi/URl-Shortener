package url_shortner

import "context"

type URLRepository interface {
	add(ctx context.Context, url URL) error
	find(ctx context.Context, id string) (*URL, error)
}

type URLConvertor interface {
	ToURL() URL
}
