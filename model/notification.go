package model

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	UserID      uint      `json:"user_id"`
	User        User      `json:"user"`
	ComplaintID uint      `json:"complaint_id"`
	Complaint   Complaint `json:"complaint"`
}
