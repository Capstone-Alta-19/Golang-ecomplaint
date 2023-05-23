package controller

import (
	"capstone/middleware"
	"capstone/model/payload"
	"capstone/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	_, err := middleware.ExtractTokenAdminId(c, "")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	users, err := usecase.GetListUsers()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

// create user
func RegisterUserController(c echo.Context) error {
	payload := payload.CreateUserRequest{}
	c.Bind(&payload)
	// validasi request body
	if err := c.Validate(payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "Invalid Request Payload",
			"errorDescription": err.Error(),
		})
	}

	if payload.Password != payload.ConfirmPassword {
		return echo.NewHTTPError(http.StatusBadRequest, "Password Not Match")
	}

	user, err := usecase.CreateUser(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}
