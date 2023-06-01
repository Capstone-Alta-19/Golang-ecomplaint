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
	PhotoURL    string   `json:"photo_url"`
	VideoURL    string   `json:"video_url"`
	IsPublic    bool     `json:"is_public"`
	Feedback    string   `json:"feedback"`
}

type Feedback struct {
	gorm.Model
	Description string `json:"description"`
}
