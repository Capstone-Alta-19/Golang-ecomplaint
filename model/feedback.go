package model

import "gorm.io/gorm"

type Feedback struct {
	gorm.Model
	ComplaintID uint   `json:"complaint_id"`
	Description string `json:"description"`
}
