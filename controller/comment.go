package controller

import (
	"capstone/middleware"
	"capstone/model/payload"
	"capstone/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateCommentController(c echo.Context) error {
	id, _ := middleware.ExtractTokenUserId(c)

	req := &payload.CreateCommentRequest{}
	c.Bind(&req)

	_, err := usecase.CreateComment(id, req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed To Create Comment",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Created Comment",
	})
}

func GetCommentController(c echo.Context) error {
	id, err := middleware.ExtractTokenUserId(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to extract user ID from token",
			"error":   err.Error(),
		})
	}

	comments, err := usecase.GetComments(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to get comments",
			"error":   err.Error(),
		})
	}

	response := map[string]interface{}{
		"message":  "Success",
		"comments": comments,
	}

	return c.JSON(http.StatusOK, response)

}
