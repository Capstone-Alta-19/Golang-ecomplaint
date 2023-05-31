package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Name     string `json:"name"`
	Role     string `json:"role"`
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
