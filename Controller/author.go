package Controller

import (
	"BearLibrary/Config"
	"BearLibrary/Helper"
	model "BearLibrary/Models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAuthors(c echo.Context) error {
	var authors []model.Author
	err := Config.DB.Find(&authors).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
      "messages": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"authors":  authors,
	})
}

func GetAuthor(c echo.Context) error {
	author := model.Author{}
	c.Bind(&author)
	id, _ := strconv.Atoi(c.Param("id"))
	err := Config.DB.Find(&author, id).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get author data",
		"author":   author,
	})
}

func AddAuthor(c echo.Context) error {
	authCheck := Helper.GetClaims(c.Request().Header.Get("Authorization"))
	if authCheck != "" {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "unauthorized",
		})
	}
	author := model.Author{}
	err := c.Bind(&author)
	if err != nil {
		return err
	}
	err = Config.DB.Save(&author).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create author",
		"author":  author,
	})
}

func UpdateAuthor(c echo.Context) error {
	authCheck := Helper.GetClaims(c.Request().Header.Get("Authorization"))
	if authCheck != "" {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "unauthorized",
		})
	}
	author := model.Author{}
	c.Bind(&author)
	id, _ := strconv.Atoi(c.Param("id"))
	err := Config.DB.Model(author).Where("id = ?", id).Updates(author).Error
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "failed to update",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update author",
	})
}

//TODO: soft-delete for author
