package Controller

import (
	"BearLibrary/Config"
	model "BearLibrary/Models"
	"github.com/labstack/echo/v4"
)

var series []model.Series

func GetSeries(c echo.Context) error {
	err := Config.DB.Find(&series)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": err,
		})
	}
	return c.JSON(200, map[string]interface{}{
		"message": "Success",
		"series":  series,
	})
}
