package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HeartBeat(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
