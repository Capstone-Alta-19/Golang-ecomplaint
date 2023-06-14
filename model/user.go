package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	DateBirth time.Time `json:"date_birth"`
	Role      string    `json:"role"`
	Token     string    `gorm:"-"`
}
