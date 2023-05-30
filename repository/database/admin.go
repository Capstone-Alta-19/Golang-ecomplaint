package database

import (
	"capstone/config"
	"capstone/model"
)

func CreateAdmin(admin *model.Admin) error {
	if err := config.DB.Create(admin).Error; err != nil {
		return err
	}
	return nil
}

func GetAdminByUsername(username string) (*model.Admin, error) {
	var admin model.Admin
	err := config.DB.Where("username = ?", username).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}
