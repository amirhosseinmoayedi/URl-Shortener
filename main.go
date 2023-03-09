package main

import (
	"log"
	"url-shortener/internal/cache"
	"url-shortener/internal/http"
	url_shortner "url-shortener/internal/url-shortner"
)

func main() {
	// init repo
	repo := cache.NewInMemoryCacheURLRepository()
	// init service
	service, err := url_shortner.NewService(repo)
	if err != nil {
		log.Fatal(err)
	}
	// init handler
	var handler *url_shortner.Handler
	handler, err = url_shortner.NewHandler(*service)
	if err != nil {
		log.Fatal(err)
	}
	// init router
	var router *http.Router
	router, err = http.NewRouter(*handler)
	// serve
	router.Serve()
}
