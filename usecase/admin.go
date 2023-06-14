package usecase

import (
	"capstone/middleware"
	"capstone/model"
	"capstone/model/payload"
	"capstone/repository/database"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateAdmin(req payload.AddAdminRequest) (*model.Admin, error) {
	_, err := database.GetAdminByUsername(req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
			if err != nil {
				return nil, err
			}
			newAdmin := model.Admin{
				Name:     req.Name,
				Role:     req.Role,
				Username: req.Username,
				Password: string(passwordHash),
			}
			err = database.CreateAdmin(&newAdmin)
			if err != nil {
				return nil, err
			}

			return &newAdmin, nil
		}
		return nil, err
	}
	return nil, errors.New("username already used")
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

func GetAdminByID(adminID uint) (*payload.GetAdminProfileResponse, error) {
	admin, err := database.GetAdminByID(adminID)
	if err != nil {
		return nil, err
	}

	resp := payload.GetAdminProfileResponse{
		ID:       admin.ID,
		Name:     admin.Name,
		Username: admin.Username,
	}

	return &resp, nil
}

func UpdateAdminByID(adminID uint, req payload.UpdateAdminRequest) error {
	admin, err := database.GetAdminByID(adminID)
	if err != nil {
		return err
	}
	if req.Name != "" {
		admin.Name = req.Name
	}
	if req.Username != "" {
		admin.Username = req.Username
	}

	if req.OldPassword != admin.Password {
		return errors.New("wrong old password")
	} else if req.NewPassword != req.ConfirmPassword {
		return errors.New("new password and confirm password not match")
	}

	admin.Password = req.NewPassword

	err = database.UpdateAdmin(admin)
	if err != nil {
		return err
	}

	return nil
}
