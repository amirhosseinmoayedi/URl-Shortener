package http

import (
	"errors"
	"fmt"
	"github.com/amirhosseinmoayedi/URl-Shortener/internal/http/handlers"
	"github.com/amirhosseinmoayedi/URl-Shortener/internal/log"
	url_shortner "github.com/amirhosseinmoayedi/URl-Shortener/internal/url-shortner"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	Port    string
	Path    string
	handler url_shortner.Handler
}

var routerPort = ""
var routerPath = ""

func NewRouter(handler url_shortner.Handler) (*Router, error) {
	if routerPort == "" {
		log.Logger.WithField("handler", handler).Info("setting port to default: 8080")
		routerPort = "8080"
	}
	if routerPath == "" {
		log.Logger.WithField("handler", handler).Info("setting port to default: 8080")
		routerPath = "localhost"
	}
	if handler == (url_shortner.Handler{}) {
		err := errors.New("handler cant be empty")
		log.Logger.WithField("handler", handler).Error(err)
		return nil, err
	}
	return &Router{
		Port:    routerPort,
		Path:    routerPath,
		handler: handler,
	}, nil
}

func (rc *Router) Serve() {
	e := rc.startRouter()

	address := fmt.Sprintf("%v:%v", rc.Path, rc.Port)

	if err := e.Start(address); err != nil {
		log.Logger.WithFields(map[string]interface{}{"address": address, "router": rc}).Fatal("cant start the server", err)
	}
}

func (rc *Router) startRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "x-CSRF-Token"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Link"},
		MaxAge:           300,
	}))
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.Logger.WithFields(map[string]interface{}{
				"URI":    values.URI,
				"status": values.Status,
			}).Info("request")
			return nil
		},
	}))

	e.GET("/health-check/", handlers.HeartBeat)

	e.POST("/shorten-url/", rc.handler.ShortenUrl)
	e.GET("/shorted-url/:uuid", rc.handler.RedirectToOrigin)

	return e
}
