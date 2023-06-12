package model

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	UserID      uint `json:"user_id"`
	ComplaintID uint `json:"complaint_id"`
}
