package controller

import (
	"capstone/constant"
	"capstone/middleware"
	"capstone/model/payload"
	"capstone/usecase"
	"net/http"

	"github.com/labstack/echo"
)

func AddAdminController(c echo.Context) error {
	role, _, err := middleware.ExtractTokenAdminId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Only Super Admin Can Access This Feature")
	}

	if role != constant.SuperAdmin {
		return echo.NewHTTPError(http.StatusUnauthorized, "Only Super Admin Can Access This Feature")
	}

	req := payload.AddAdminRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
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
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Login",
		"admin":   admin,
	})
}

func GetAdminController(c echo.Context) error {
	role, adminID, err := middleware.ExtractTokenAdminId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Only Admin Can Access This Feature")
	}

	if role != constant.SuperAdmin && role != constant.Admin {
		return echo.NewHTTPError(http.StatusUnauthorized, "Only Admin Can Access This Feature")
	}

	admin, err := usecase.GetAdminByID(adminID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get Admin",
		"admin":   admin,
	})
}

func UpdateAdminController(c echo.Context) error {
	role, adminID, err := middleware.ExtractTokenAdminId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Only Admin Can Access This Feature")
	}

	if role != constant.SuperAdmin && role != constant.Admin {
		return echo.NewHTTPError(http.StatusUnauthorized, "Only Admin Can Access This Feature")
	}

	req := payload.UpdateAdminRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = usecase.UpdateAdminByID(adminID, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Success Updated Admin")
}
