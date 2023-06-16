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
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid sort query,must be 'asc' or 'desc'")
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
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid status query, must be 'All' or 'Pending' or 'Proccess' or 'Resolved'")
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

func GetAllComplaintsController(c echo.Context) error {
	role, _, err := middleware.ExtractTokenAdminId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if role != constant.Admin && role != constant.SuperAdmin {
		return echo.NewHTTPError(http.StatusBadRequest, "You are not authorized")
	}

	sortBy := c.QueryParam("sort")
	if sortBy != constant.Ascending && sortBy != constant.Descending {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid sort query, must be 'asc' or 'desc'")
	}

	typeSort := c.QueryParam("type")
	if typeSort != constant.Complaint && typeSort != constant.Aspiration && typeSort != constant.StatusAll {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid type query, must be 'Complaint' or 'Aspiration' or 'All'")
	}

	search := c.QueryParam("search")

	page := c.QueryParam("page")
	if page == "" {
		page = "1"
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	limit := c.QueryParam("limit")
	if limit == "" {
		limit = "10"
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	complaints, err := usecase.GetAllComplaints(sortBy, typeSort, search, limitInt, pageInt)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "Success",
		"complaints": complaints,
	})
}

func CreateFeedbackByComplaintIDController(c echo.Context) error {
	role, _, err := middleware.ExtractTokenAdminId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if role != constant.Admin && role != constant.SuperAdmin {
		return echo.NewHTTPError(http.StatusBadRequest, "You are not authorized")
	}

	complaintID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	req := &payload.CreateFeedbackRequest{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	_, err = usecase.CreateFeedbackByComplaintID(req, uint(complaintID))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}

func UpdateComplaintController(c echo.Context) error {
	role, _, err := middleware.ExtractTokenAdminId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if role != constant.Admin && role != constant.SuperAdmin {
		return echo.NewHTTPError(http.StatusBadRequest, "You are not authorized")
	}

	complaintID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	req := &payload.UpdateComplaintRequest{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	_, err = usecase.UpdateComplaintByID(req, uint(complaintID))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}

func GetComplaintByIDController(c echo.Context) error {
	userID, err := middleware.ExtractTokenUserId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	complaint, err := usecase.GetUserComplaintID(uint(id), userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Success",
		"complaint": complaint,
	})
}

func AdminDeleteComplaintByIDController(c echo.Context) error {
	role, _, err := middleware.ExtractTokenAdminId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if role != constant.Admin && role != constant.SuperAdmin {
		return echo.NewHTTPError(http.StatusBadRequest, "You are not authorized")
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = usecase.AdminDeleteComplaintByID(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}

func PinnedComplaintByIDController(c echo.Context) error {
	userID, err := middleware.ExtractTokenUserId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	complaintID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = usecase.PinnedComplaint(userID, uint(complaintID))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}

func UnpinnedComplaintByIDController(c echo.Context) error {
	userID, err := middleware.ExtractTokenUserId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	complaintID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = usecase.UnpinnedComplaint(userID, uint(complaintID))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}

func GetPinnedComplaintController(c echo.Context) error {
	userID, err := middleware.ExtractTokenUserId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	complaints, err := usecase.GetPinnedComplaint(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "Success",
		"complaints": complaints,
	})
}
