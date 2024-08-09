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

	//Authors
	e.GET("/authors", Controller.GetAuthors, middleware.JWT([]byte("secret")))
	e.GET("/authors/:id", Controller.GetAuthor)
	e.POST("/authors", Controller.AddAuthor, middleware.JWT([]byte("secret")))
	e.PUT("/authors/:id", Controller.UpdateAuthor, middleware.JWT([]byte("secret")))

	//Series
	e.GET("/series", Controller.GetSeries)
	e.GET("/series/:id", Controller.GetSeriesByAuthor)
	e.POST("series", Controller.AddSeries)

	//Users
	e.POST("/users", Controller.Login)
	//healthcheck
	e.GET("/check", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "ok",
		})
	})
	return e
}
