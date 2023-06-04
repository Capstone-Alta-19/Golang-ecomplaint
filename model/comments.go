package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	NewsID      string `json:"news_id"`
	Description string `json:"description"`
}
