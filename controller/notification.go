package controller

import (
	"capstone/constant"
	"capstone/middleware"
	"capstone/usecase"
	"net/http"

	"github.com/labstack/echo"
)

func GetNotificationController(c echo.Context) error {
	role, _, err := middleware.ExtractTokenAdminId(c)
	if err != nil {
		return err
	}
	if role != constant.SuperAdmin && role != constant.Admin {
		return echo.NewHTTPError(http.StatusUnauthorized, "Only Admin Can Access This Feature")
	}

	notifications, err := usecase.GetNotification()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "Success Get Notification",
		"notifications": notifications,
	})
}
