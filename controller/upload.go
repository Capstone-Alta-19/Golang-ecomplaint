package controller

import (
	"capstone/constant"
	"capstone/utils"
	"net/http"

	"github.com/labstack/echo"
)

func UploadFileController(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	url, err := utils.UploadCloudinary(file, constant.AppName)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"url":     url,
	})
}
