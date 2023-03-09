package middlewares

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func HeartBeatMiddleware(path string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if (c.Request().Method == "GET" || c.Request().Method == "HEAD") && strings.EqualFold(path, c.Request().URL.Path) {
				return c.String(http.StatusOK, "pong")
			}
			return next(c)
		}
	}
}
