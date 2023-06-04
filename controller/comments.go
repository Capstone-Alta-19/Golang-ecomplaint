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

func GetCommentController(c echo.Context) error {
	comment, e := usecase.GetListComment()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"Comment": comment,
	})
}

func CreateCommentController(c echo.Context) error {
	role, _, err := middleware.ExtractTokenAdminId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	if role != "admin" {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	payload := payload.CreateComment{}
	c.Bind(&payload)
	if err := c.Validate(payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request payload",
		})
	}

	comment, err := usecase.CreateComment(payload.NewsID, payload.Description)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success add comment",
		"comment": comment,
	})
}

func DeleteCommentController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := usecase.DeleteComment(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete comment",
			"errorDescription": err,
			"errorMessage":     "Sorry, the comment cannot be deleted",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete comment",
	})
}

func UpdateCommentController(c echo.Context) error {
	role, _, err := middleware.ExtractTokenAdminId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	if role != "admin" {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	comment := model.Comment{}
	c.Bind(&comment)
	comment.ID = uint(id)

	if err := usecase.UpdateComment(&comment); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update comment",
			"errorDescription": err,
			"errorMessage":     "Sorry, the comment cannot be changed",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update comment",
	})
}
