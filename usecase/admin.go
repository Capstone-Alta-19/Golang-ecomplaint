package usecase

import (
	"capstone/middleware"
	"capstone/model"
	"capstone/model/payload"
	"capstone/repository/database"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateAdmin(req payload.AddAdminRequest) (*model.Admin, error) {
	_, err := database.GetAdminByUsername(req.Username)
	if err == nil {
		return nil, errors.New("username already used")
	}

	newAdmin := model.Admin{
		Name:     req.Name,
		Role:     req.Role,
		Username: req.Username,
		Password: req.Password,
	}

	err = database.CreateAdmin(&newAdmin)
	if err != nil {
		return nil, err
	}
	return &newAdmin, nil
}

func LoginAdmin(req payload.LoginAdminRequest) (*model.Admin, error) {
	admin, err := database.GetAdminByUsername(req.Username)
	if err != nil {
		return nil, err
	}
	if admin.Password != req.Password {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Wrong Password")
	}
	// generate jwt
	token, err := middleware.CreateToken(admin.ID, admin.Role)
	if err != nil {
		fmt.Println("LoginUser: Error generating token")
		return nil, err
	}
	admin.Token = token
	return admin, nil
}
