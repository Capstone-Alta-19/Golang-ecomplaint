package utils

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {

			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				validationErrors = append(validationErrors, err.Field()+" is required")
			case "email":
				validationErrors = append(validationErrors, err.Field()+" is not valid email")
			case "oneof":
				validationErrors = append(validationErrors, err.Field()+" must be one of "+err.Param())
			default:
				validationErrors = append(validationErrors, err.Error())
			}
		}
		return echo.NewHTTPError(http.StatusBadRequest, validationErrors)
	}
	return nil
}
