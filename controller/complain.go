package controller

import (
	"capstone/constant"
	"capstone/middleware"
	"capstone/model/payload"
	"capstone/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func CreateComplaintController(c echo.Context) error {
	id, err := middleware.ExtractTokenUserId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	req := &payload.CreateComplaintRequest{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	_, err = usecase.CreateComplaint(id, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed To Create Complaint",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Created Complaint",
	})
}

func GetComplaintController(c echo.Context) error {
	id, err := middleware.ExtractTokenUserId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to extract user ID from token",
			"error":   err.Error(),
		})
	}

	complaints, err := usecase.GetComplaints(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to get complaints",
			"error":   err.Error(),
		})
	}

	response := map[string]interface{}{
		"message":    "Success",
		"complaints": complaints,
	}

	return c.JSON(http.StatusOK, response)

}

func GetComplaintByIDController(c echo.Context) error {
	role, _, err := middleware.ExtractTokenAdminId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if role != constant.Admin && role != constant.SuperAdmin {
		return echo.NewHTTPError(http.StatusBadRequest, "You are not authorized")
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	complaint, err := usecase.GetComplaintByID(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Success",
		"complaint": complaint,
	})
}

func GetFeedbackController(c echo.Context) error {
	complaintID := c.Param("complaintID")

	feedback, err := usecase.GetFeedback(complaintID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to get feedback").SetInternal(err)
	}

	response := map[string]interface{}{
		"message":  "Success",
		"feedback": feedback,
	}

	return c.JSON(http.StatusOK, response)
}
