package controller

import (
	"capstone/middleware"
	"capstone/model/payload"
	"capstone/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddAdminController(c echo.Context) error {
	_, err := middleware.ExtractTokenAdminId(c, "Super Admin")
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Only Super Admin Can Access This Feature")
	}
	req := payload.AddAdminRequest{}
	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Request")
	}

	_, err = usecase.CreateAdmin(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "Success Added Admin")
}

func LoginAdminController(c echo.Context) error {
	req := payload.LoginAdminRequest{}
	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Request Payload ")
	}

	admin, err := usecase.LoginAdmin(req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Login",
		"admin":   admin,
	})
}
