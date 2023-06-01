package model

import (
	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	NewsName    string `json:"news_name"`
	Description string `json:"description"`
}
