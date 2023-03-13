package main

import (
	"github.com/amirhosseinmoayedi/URl-Shortener/internal/cache"
	"github.com/amirhosseinmoayedi/URl-Shortener/internal/http"
	log "github.com/amirhosseinmoayedi/URl-Shortener/internal/log"
	urlshortner "github.com/amirhosseinmoayedi/URl-Shortener/internal/url-shortner"
)

func init() {
	log.InitLogger()
}

func main() {
	repo := cache.NewInMemoryCacheURLRepository()

	service, err := url_shortner.NewService(repo)
	if err != nil {
		log.Logger.Fatal(err)
	}
	var handler *url_shortner.Handler
	handler, err = url_shortner.NewHandler(*service)
	if err != nil {
		log.Logger.Fatal(err)
	}

	var router *http.Router
	router, err = http.NewRouter(*handler)
	if err != nil {
		log.Logger.Fatal(err)
	}
	router.Serve()
}
