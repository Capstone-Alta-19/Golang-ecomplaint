package model

import "gorm.io/gorm"

type Feedback struct {
	gorm.Model
	Description string `json:"description"`
}
