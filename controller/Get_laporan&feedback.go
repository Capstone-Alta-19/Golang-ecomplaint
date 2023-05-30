package controller

import (
	"capstone/middleware"
	"capstone/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetComplaintController(c echo.Context) error {
	id, _ := middleware.ExtractTokenUserId(c)

	complaints, err := usecase.GetComplaints(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed To Get Complaints",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "Success",
		"complaints": complaints,
	})
}
