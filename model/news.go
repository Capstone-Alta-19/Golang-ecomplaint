package model

import (
	"time"

	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	NewsName    string    `json:"news_name"`
	Description string    `json:"description"`
	CategoryID  uint      `json:"category_id"`
	Category    Category  `json:"-"`
	Time        time.Time `json:"-"`
}
