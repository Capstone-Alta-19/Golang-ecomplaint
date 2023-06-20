package controller

import (
	"capstone/constant"
	"capstone/middleware"
	"capstone/model"
	"capstone/model/payload"
	"capstone/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetNewsController(c echo.Context) error {
	news, e := usecase.GetListNews()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"News":   news,
	})
}

func CreateNewsController(c echo.Context) error {
	role, adminID, err := middleware.ExtractTokenAdminId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	if role != constant.Admin && role != constant.SuperAdmin {
		return echo.NewHTTPError(http.StatusUnauthorized, "Only Admin Can Access This Feature")
	}

	payload := payload.CreateNews{}
	c.Bind(&payload)
	if err := c.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	news, err := usecase.CreateNews(&payload, adminID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success add news",
		"news":    news,
	})
}

func DeleteNewsController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := usecase.DeleteNews(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete news",
			"errorDescription": err,
			"errorMessage":     "Sorry, the news cannot be deleted",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete news",
	})
}

func UpdateNewsController(c echo.Context) error {
	role, _, err := middleware.ExtractTokenAdminId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Only Admin Can Access This Feature")
	}

	if role != constant.Admin {
		return echo.NewHTTPError(http.StatusUnauthorized, "Only Admin Can Access This Feature")
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	news := model.News{}
	c.Bind(&news)
	news.ID = uint(id)

	if err := usecase.UpdateNews(&news); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update news",
			"errorDescription": err,
			"errorMessage":     "Sorry, the news cannot be changed",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update news",
	})
}

func GetNewsByIDController(c echo.Context) error {
	_, err := middleware.ExtractTokenUserId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Only User Can Access This Feature")
	}

	newsID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	news, err := usecase.GetNewsByID(uint(newsID))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"news":   news,
	})
}
