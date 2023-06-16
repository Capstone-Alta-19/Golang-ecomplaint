package usecase

import (
	"capstone/middleware"
	"capstone/model"
	"capstone/model/payload"
	"capstone/repository/database"
	"errors"
	"fmt"

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

func LoginAdmin(req payload.LoginAdminRequest) (*payload.LoginAdminResponse, error) {
	admin, err := database.GetAdminByUsername(req.Username)
	if err != nil {
		return nil, errors.New("username not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, errors.New("wrong password")
	}

	// generate jwt
	token, err := middleware.CreateToken(admin.ID, admin.Role)
	if err != nil {
		fmt.Println("LoginUser: Error generating token")
		return nil, err
	}
	admin.Token = token
	resp := payload.LoginAdminResponse{
		Token: token,
	}
	return &resp, nil
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

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.OldPassword))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return errors.New("old password is not correct")
	}

	if req.NewPassword != req.ConfirmPassword {
		return errors.New("new password and confirm password not match")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	admin.Password = string(passwordHash)
	err = database.UpdateAdmin(admin)
	if err != nil {
		return err
	}

	return nil
}
