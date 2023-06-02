package controller

import (
	"net/http"

	"capstone/model/payload"
	"capstone/usecase"

	"github.com/labstack/echo"
)

func LoginUserController(c echo.Context) error {
	req := payload.LoginUserRequest{}
	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user, err := usecase.LoginUser(req.UsernameOrEmail, req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"user":    user,
	})
}
