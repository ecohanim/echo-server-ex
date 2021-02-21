package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetAdminMiddlewares(g *echo.Group) {
	// this logs the server interaction
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}` + "\n",
	}))

	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// check in the DB if credentials are valid
		if username == "admin" && password == "admin" {
			return true, nil
		}
		return false, nil
	}))
}
