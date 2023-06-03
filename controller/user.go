package controller

import (
	"capstone/middleware"
	"capstone/model"
	"capstone/model/payload"
	"capstone/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

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

func UpdateUserController(c echo.Context) error {
	_, err := middleware.ExtractTokenUserId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Only User Can Access This Feature")
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	user := model.User{}
	c.Bind(&user)
	user.ID = uint(id)

	if err := usecase.UpdateUser(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update user",
			"errorDescription": err,
			"errorMessage":     "Sorry, the user cannot be changed",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
	})
}
