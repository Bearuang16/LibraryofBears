package Route

import (
	"BearLibrary/Controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func New() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/authors", Controller.GetAuthors)
	e.GET("/authors/:id", Controller.GetAuthor)
	e.POST("/authors", Controller.AddAuthor)
	e.PUT("/authors/:id", Controller.UpdateAuthor)

	e.GET("/check", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "ok",
		})
	})
	return e
}
