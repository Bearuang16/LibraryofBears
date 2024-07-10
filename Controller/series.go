package Controller

import (
	"BearLibrary/Config"
	model "BearLibrary/Models"
	"net/http"

	"github.com/labstack/echo/v4"
)

var series []model.Series

func GetSeries(c echo.Context) error {
	err := Config.DB.Find(&series)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"series":  series,
	})
}
