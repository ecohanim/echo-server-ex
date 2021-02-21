package api

import (
	"ecohanim/echo-server-ex/web-server/api/handlers"

	"github.com/labstack/echo"
)

func CookieGroup(g *echo.Group) {
	g.GET("/main", handlers.MainAdmin)
}
