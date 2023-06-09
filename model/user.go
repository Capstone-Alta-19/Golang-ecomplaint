package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	PhotoProfile string `json:"photo_profile"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	FullName     string `json:"full_name"`
	Password     string `json:"password"`
	Phone        string `json:"phone"`
	DateBirth    string `json:"date_birth"`
	Role         string `json:"role"`
	Token        string `gorm:"-"`
}
