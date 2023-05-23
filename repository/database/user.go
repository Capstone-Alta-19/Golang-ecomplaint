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

// mendapatkan data semua user
func GetUsers() (users []model.User, err error) {
	if err = config.DB.Find(&users).Error; err != nil {
		return
	}
	return
}

// mendapatkan data user dengan ID
func GetUser(id uint) (*model.User, error) {
	var user model.User
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
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

// login user
func LoginUser(user *model.User) error {
	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return err
	}
	return nil
}
