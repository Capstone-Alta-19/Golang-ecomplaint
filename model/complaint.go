package model

import (
	"gorm.io/gorm"
)

type Complaint struct {
	gorm.Model
	UserID      uint      `json:"user_id"`
	User        User      `json:"user"`
	CategoryID  uint      `json:"category_id"`
	Category    Category  `json:"-"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	PhotoURL    string    `json:"photo_url"`
	VideoURL    string    `json:"video_url"`
	IsPublic    bool      `json:"is_public"`
	Status      string    `json:"status"`
	LikesCount  uint      `json:"likes_count"`
	Feedback    Feedback  `json:"feedback"`
	Comments    []Comment `json:"comments"`
}
