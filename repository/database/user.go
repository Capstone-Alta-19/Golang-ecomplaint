package database

import (
	"capstone/config"
	"capstone/model"
)

// membuat user
func CreateUser(user *model.User) error {
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// mendapatkan data user dengan Email
func GetUserByUsernameOrEmail(UsernameOrEmail string) (*model.User, error) {
	var user model.User

	err := config.DB.Where("username = ? OR email = ?", UsernameOrEmail, UsernameOrEmail).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByUsername(Username string) (*model.User, error) {
	var user model.User

	err := config.DB.Where("username = ?", Username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(Email string) (*model.User, error) {
	var user model.User

	err := config.DB.Where("email = ?", Email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// update user
func UpdateUser(user *model.User) error {
	if err := config.DB.Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByID(id uint) (*model.User, error) {
	var user model.User
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func DeleteUser(user *model.User) error {
	if err := config.DB.Delete(user).Error; err != nil {
		return err
	}
	return nil
}
