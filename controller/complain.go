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

func GetComplaintsByCategoryIDController(c echo.Context) error {
	_, err := middleware.ExtractTokenUserId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	queryParam := c.QueryParam("sort")
	if queryParam != constant.Ascending && queryParam != constant.Descending {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid sort query")
	}

	complaints, err := usecase.GetComplaintsByCategoryID(uint(id), queryParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "Success",
		"complaints": complaints,
	})
}

func GetUserComplaintsByStatusController(c echo.Context) error {
	userID, err := middleware.ExtractTokenUserId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	status := c.QueryParam("status")
	if status != constant.StatusAll && status != constant.StatusPending && status != constant.StatusProccess && status != constant.StatusResolved {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid status query")
	}

	complaints, err := usecase.GetUserComplaintsByStatus(userID, status)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "Success",
		"complaints": complaints,
	})
}

func AdminGetComplaintByIDController(c echo.Context) error {
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

func DeleteComplaintByIDController(c echo.Context) error {
	userID, err := middleware.ExtractTokenUserId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	complaintID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = usecase.DeleteComplaintByID(userID, uint(complaintID))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}

func GetTotalComplaintsController(c echo.Context) error {
	role, _, err := middleware.ExtractTokenAdminId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if role != constant.Admin && role != constant.SuperAdmin {
		return echo.NewHTTPError(http.StatusBadRequest, "You are not authorized")
	}

	total, err := usecase.GetTotalComplaints()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"total":   total,
	})
}
