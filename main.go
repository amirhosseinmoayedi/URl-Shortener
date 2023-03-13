package main

import (
	service2 "github.com/amirhosseinmoayedi/URl-Shortener/internal/domain/service"
	"github.com/amirhosseinmoayedi/URl-Shortener/internal/infrastructure/db/memory"
	"github.com/amirhosseinmoayedi/URl-Shortener/internal/interface/http"
	urlshortner "github.com/amirhosseinmoayedi/URl-Shortener/internal/interface/http/v1"
)

func main() {
	repo := memory.NewInMemoryCacheURLRepository()
	service := service2.NewService(repo)
	handler := urlshortner.NewHandler(service)

	router := http.NewRouter(handler)
	router.Serve()
}
