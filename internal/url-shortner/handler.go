package url_shortner

import (
	"errors"
	"github.com/amirhosseinmoayedi/URl-Shortener/internal/log"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

type urlShortenResponse struct {
	URL       string    `json:"url"`
	Shorten   string    `json:"shorten"`
	CreatedAt time.Time `json:"created_at"`
}

type urlShortenRequest struct {
	webSiteDomain string `query:"domain" validate:"required"`
}
type urlRedirectRequest struct {
	path string `param:"path" validate:"required"`
}

func (h Handler) RedirectToOrigin(ctx echo.Context) error {
	var request urlRedirectRequest
	if err := ctx.Bind(&request); err != nil {
		log.Logger.WithFields(map[string]interface{}{"request": *ctx.Request(), "err": err}).Error("can't bind to urlRedirectRequest")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := ctx.Validate(request); err != nil {
		log.Logger.WithFields(map[string]interface{}{"request": *ctx.Request(), "err": err}).Error("validation error for urlRedirectRequest")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	bctx := ctx.Request().Context()
	url, err := h.svc.GetShortenLink(bctx, request.path)
	if err != nil {
		if errors.Is(err, URLNotFound) {
			log.Logger.WithFields(map[string]interface{}{"request": *ctx.Request(), "err": err, "url": url}).Error(err)
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		} else {
			log.Logger.WithFields(map[string]interface{}{"request": *ctx.Request(), "err": err, "url": url}).Error("can't retrieve the url for uuid")
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	responsePayload := urlShortenResponse{
		URL:       url.Original,
		Shorten:   url.Path,
		CreatedAt: url.CreatedAt,
	}
	log.Logger.WithFields(map[string]interface{}{"ctx": ctx, "url": url, "response": responsePayload}).Info("short-link created")
	return ctx.Redirect(http.StatusMovedPermanently, responsePayload.Shorten)
}

func (h Handler) ShortenUrl(ctx echo.Context) error {
	var request urlShortenRequest
	err := ctx.Bind(&request)
	if err != nil {
		log.Logger.WithFields(map[string]interface{}{"request": *ctx.Request(), "err": err}).Error("can't bind the request for shorten url")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err = ctx.Validate(request); err != nil {
		log.Logger.WithFields(map[string]interface{}{"request": *ctx.Request(), "err": err}).Error("validation error for urlShortenRequest")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	bctx := ctx.Request().Context()
	url := NewUrl(request.webSiteDomain)
	if err = h.svc.ShortenLink(bctx, url); err != nil {
		log.Logger.WithFields(map[string]interface{}{"request": *ctx.Request(), "err": err, "url": url}).Error("can't create a shorted url")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	responsePayload := urlShortenResponse{
		URL:       url.Original,
		Shorten:   url.Path,
		CreatedAt: url.CreatedAt,
	}
	log.Logger.WithFields(map[string]interface{}{"request": *ctx.Request(), "err": err, "url": url, "response": responsePayload}).Info("response to shorted url request")
	return ctx.JSONPretty(http.StatusCreated, responsePayload, "\t")
}
