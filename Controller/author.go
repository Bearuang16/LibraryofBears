package Controller

import (
	"BearLibrary/Config"
	model "BearLibrary/Models"
	"github.com/labstack/echo/v4"
	"strconv"
)

var authors []model.Author

func GetAuthors(c echo.Context) error {
	err := Config.DB.Find(&authors)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error,
		})
	}
	return c.JSON(200, map[string]interface{}{
		"messages": "success",
		"authors":  authors,
	})
}

func GetAuthorController(c echo.Context) error {
	author := model.Author{}
	c.Bind(&author)
	id, _ := strconv.Atoi(c.Param("id"))
	err := Config.DB.Find(&author, id).Error
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(200, map[string]interface{}{
		"messages": "success get author data",
		"author":   author,
	})
}

func AddAuthor(c echo.Context) error {
	author := model.Author{}
	err := c.Bind(&author)
	if err != nil {
		return err
	}
	err = Config.DB.Save(&author).Error
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(200, map[string]interface{}{
		"message": "success create author",
		"author":  author,
	})
}
