package controller

import (
	"net/http"
	"project_structure/usecase"
	"strconv"

	"github.com/labstack/echo"
)

func GetReportController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	report, err := usecase.GetReport(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error retrieving report",
			"error":   err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"report": report,
	})
}
