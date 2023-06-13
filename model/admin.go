package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Name     string `json:"name"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
