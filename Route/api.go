package Route

import (
	config "BearLibrary/Config"
	"BearLibrary/Controller"
	"github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func New() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())

	//Authors
	e.GET("/authors", Controller.GetAuthors)
	e.GET("/authors/:id", Controller.GetAuthor)
	e.POST("/authors", Controller.AddAuthor, echojwt.JWT([]byte(config.Secret)))
	e.PUT("/authors/:id", Controller.UpdateAuthor, echojwt.JWT([]byte(config.Secret)))

	//Series
	e.GET("/series", Controller.GetSeries)
	e.GET("/series/:id", Controller.GetSeriesByAuthor, echojwt.JWT([]byte(config.Secret)))
	e.POST("/series", Controller.AddSeries)

	//Users
	e.POST("/users", Controller.Login)
	//healthcheck
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "ok",
		})
	})
	return e
}
