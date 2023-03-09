package http

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"url-shortener/internal/http/middlewares"
	url_shortner "url-shortener/internal/url-shortner"
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
		log.Print("setting port to default: 8080")
		routerPort = "8080"
	}
	if routerPath == "" {
		log.Print("setting path to default localhost")
		routerPath = "localhost"
	}
	if handler == (url_shortner.Handler{}) {
		return nil, errors.New("handler cant be empty")
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
		log.Fatal("cant start the server", err)
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
	e.Use(middlewares.HeartBeatMiddleware("/ping"))

	e.POST("/shorten-url/", rc.handler.ShortenUrl)
	e.GET("/shorted-url/:uuid", rc.handler.RedirectToOrigin)

	return e
}
