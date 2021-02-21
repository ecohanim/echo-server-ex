package api

import (
	"ecohanim/echo-server-ex/web-server/api/handlers"

	"github.com/labstack/echo"
)

func JwtGroup(g *echo.Group) {
	g.GET("/main", handlers.MainAdmin)
}
