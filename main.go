package main

import (
	"github.com/amirhosseinmoayedi/URl-Shortener/internal/domain/service"
	"github.com/amirhosseinmoayedi/URl-Shortener/internal/infrastructure/db/postgres"
	"github.com/amirhosseinmoayedi/URl-Shortener/internal/interface/http"
	urlshortner "github.com/amirhosseinmoayedi/URl-Shortener/internal/interface/http/v1"
	"github.com/amirhosseinmoayedi/URl-Shortener/internal/log"
	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	//repo := memory.NewURLRepositoryInMemory()
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tehran"
	db, err := gorm.Open(gormPostgres.New(gormPostgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		log.Logger.WithFields(map[string]interface{}{"dsn": dsn, "err": err}).Fatal("can't open the postgres connection")
	}
	err = db.AutoMigrate(&postgres.URLPostgres{})
	if err != nil {
		log.Logger.WithField("err", err).Fatal("can't make database migrate ")
	}

	repo := postgres.NewPostgresURLRepository(db)
	srv := service.NewService(repo)
	handler := urlshortner.NewHandler(srv)

	router := http.NewServer(handler)
	router.Serve()
}
