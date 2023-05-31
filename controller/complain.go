package controller

import (
	"capstone/middleware"
	"capstone/model/payload"
	"capstone/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateComplaintController(c echo.Context) error {
	id, _ := middleware.ExtractTokenUserId(c)

	req := &payload.CreateComplaintRequest{}
	c.Bind(&req)

	_, err := usecase.CreateComplaint(id, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed To Create Complaint",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Created Complaint",
	})
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
}	
