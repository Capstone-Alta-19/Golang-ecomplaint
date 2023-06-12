package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	ComplaintID uint      `json:"complaint_id"`
	Complaint   Complaint `json:"-"`
	UserID      uint      `json:"user_id"`
	User        User      `json:"user"`
	Description string    `json:"description"`
}
