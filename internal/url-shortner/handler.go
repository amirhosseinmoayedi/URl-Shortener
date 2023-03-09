package url_shortner

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"time"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) (*Handler, error) {
	if svc == (Service{}) {
		return nil, errors.New("service cant be empty struct")
	}
	return &Handler{
		svc: svc,
	}, nil
}

type urlShortenResponse struct {
	URL       string    `json:"url"`
	Shorten   string    `json:"shorten"`
	CreatedAt time.Time `json:"created_at"`
}

func newURLShortenResponse(url URL) *urlShortenResponse {
	return &urlShortenResponse{
		URL:       url.Original,
		Shorten:   url.Path,
		CreatedAt: url.CreatedAt,
	}
}

func (h Handler) RedirectToOrigin(ctx echo.Context) error {
	uid := ctx.Param("uuid")
	if uid == "" {
		err := errors.New("uuid can't be empty")
		log.Info(err)
		return ctx.JSON(http.StatusBadRequest, err)
	}

	bctx := context.Background()
	url, err := h.svc.GetShortenLink(bctx, uid)
	if err != nil {
		if !errors.Is(err, URLNotFound) {
			log.Info()
			return ctx.JSON(http.StatusNotFound, err)
		} else {
			err = fmt.Errorf("can't retrevive the url for uuid %v: %w", uid, err)
			log.Error(err)
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}

	responsePayload := newURLShortenResponse(*url)
	return ctx.Redirect(http.StatusFound, responsePayload.Shorten)
}

func (h Handler) ShortenUrl(ctx echo.Context) error {
	givenURL := ctx.QueryParam("givenURL")
	if givenURL == "" {
		err := errors.New("givenURL params cant be empty")
		log.Info(err)
		return ctx.JSON(http.StatusBadRequest, err)
	}

	bctx := context.Background()
	url := NewUrl(givenURL)
	if err := h.svc.ShortenLink(bctx, url); err != nil {
		log.Info(err)
		return ctx.JSON(http.StatusBadRequest, err)
	}

	responsePayload := newURLShortenResponse(*url)

	return ctx.JSONPretty(http.StatusCreated, responsePayload, "\t")
}
