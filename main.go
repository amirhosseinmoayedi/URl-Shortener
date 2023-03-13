package main

import (
	"github.com/amirhosseinmoayedi/URl-Shortener/internal/cache"
	"github.com/amirhosseinmoayedi/URl-Shortener/internal/http"
	urlshortner "github.com/amirhosseinmoayedi/URl-Shortener/internal/url-shortner"
)

func main() {
	repo := cache.NewInMemoryCacheURLRepository()
	service := urlshortner.NewService(repo)
	handler := urlshortner.NewHandler(service)

	router := http.NewRouter(handler)
	router.Serve()
}
