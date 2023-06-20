package model

import (
	"time"

	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	NewsName    string    `json:"news_name"`
	PhotoURL    string    `json:"photo_url"`
	Description string    `json:"description"`
	CategoryID  uint      `json:"category_id"`
	Category    Category  `json:"-"`
	AdminID     uint      `json:"admin_id"`
	Admin       Admin     `json:"-"`
	Time        time.Time `json:"-"`
}
