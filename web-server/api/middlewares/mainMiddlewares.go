package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetMainMiddlewares(e *echo.Echo) {
	// e.Use(middleware.Static("/tmp/static")) // simple way of serving static dir
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root: "/tmp/static",
	}))

	e.Use(serverHeader)

}

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "CHQ/1.0")
		c.Response().Header().Set("NotRealHeader", "thisHaveNoMeaning")
		return next(c)
	}
}

