package model

import "gorm.io/gorm"

type Complaint struct {
	gorm.Model
	UserID      uint     `json:"user_id"`
	User        User     `json:"-"`
	CategoryID  uint     `json:"category_id"`
	Category    Category `json:"-"`
	Type        string   `json:"type"`
	Description string   `json:"description"`
	Feedback    string   `json:"feedback"`
}
