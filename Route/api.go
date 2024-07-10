package Route

import "github.com/labstack/echo/v4"

func New() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error { return nil })
	return e
}
