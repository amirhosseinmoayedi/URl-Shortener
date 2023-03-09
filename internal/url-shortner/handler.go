package url_shortner

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
	"url-shortener/internal/log"
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
	path := ctx.Param("path")
	if path == "" {
		err := errors.New("uuid can't be empty")
		log.Logger.WithField("ctx", ctx).Debug(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	bctx := context.Background()
	url, err := h.svc.GetShortenLink(bctx, path)
	if err != nil {
		if !errors.Is(err, URLNotFound) {
			log.Logger.WithFields(map[string]interface{}{"ctx": ctx, "url": url}).Debug(err)
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		} else {
			err = fmt.Errorf("can't retrevive the url for uuid %v: %w", path, err)
			log.Logger.WithFields(map[string]interface{}{"ctx": ctx, "url": url}).Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	responsePayload := newURLShortenResponse(*url)
	log.Logger.WithFields(map[string]interface{}{"ctx": ctx, "url": url, "response": responsePayload}).Info("short-link created")
	return ctx.Redirect(http.StatusMovedPermanently, responsePayload.Shorten)
}

func (h Handler) ShortenUrl(ctx echo.Context) error {
	givenURL := ctx.QueryParam("givenURL")
	if givenURL == "" {
		err := errors.New("givenURL params cant be empty")
		log.Logger.WithFields(map[string]interface{}{"ctx": ctx, "givenURL": givenURL}).Debug(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	bctx := context.Background()
	url := NewUrl(givenURL)
	if err := h.svc.ShortenLink(bctx, url); err != nil {
		log.Logger.WithFields(map[string]interface{}{"ctx": ctx, "givenURL": givenURL, "url": url}).Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	responsePayload := newURLShortenResponse(*url)
	log.Logger.WithFields(map[string]interface{}{"ctx": ctx, "url": url, "response": responsePayload}).Info("redirected")
	return ctx.JSONPretty(http.StatusCreated, responsePayload, "\t")
}
