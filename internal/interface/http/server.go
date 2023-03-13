package http

import (
	"fmt"
	urlshortner "github.com/amirhosseinmoayedi/URl-Shortener/internal/interface/http/v1"
	"github.com/amirhosseinmoayedi/URl-Shortener/internal/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Port    string
	Path    string
	handler *urlshortner.Handler
}

var routerPort = ""
var routerPath = ""

func NewServer(handler *urlshortner.Handler) *Server {
	if routerPort == "" {
		log.Logger.Info("setting port to default: 8080")
		routerPort = "8080"
	}
	if routerPath == "" {
		log.Logger.Info("setting path to default: localhost")
		routerPath = "localhost"
	}
	return &Server{
		Port:    routerPort,
		Path:    routerPath,
		handler: handler,
	}
}

func (rc *Server) Serve() {
	e := rc.initiate()

	address := fmt.Sprintf("%v:%v", rc.Path, rc.Port)

	if err := e.Start(address); err != nil {
		log.Logger.WithFields(map[string]interface{}{"address": address, "router": rc}).Fatal("cant start the server", err)
	}
}

func (rc *Server) initiate() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "x-CSRF-Token"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Link"},
		MaxAge:           300,
	}))
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogRemoteIP:  true,
		LogHost:      true,
		LogMethod:    true,
		LogURI:       true,
		LogUserAgent: true,
		LogStatus:    true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.Logger.WithFields(map[string]interface{}{
				"URI":        values.URI,
				"status":     values.Status,
				"remote_ip":  values.RemoteIP,
				"host":       values.Host,
				"method":     values.Method,
				"user_agent": values.UserAgent,
			}).Info("request")
			return nil
		},
	}))

	e.GET("/health-check/", urlshortner.HeartBeat)

	e.POST("/shorten-url/", rc.handler.ShortenUrl)
	e.GET("/shorted-url/:path/", rc.handler.RedirectToOrigin)

	return e
}
