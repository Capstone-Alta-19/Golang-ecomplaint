package payload

import (
	"gorm.io/gorm"
)

type CreateNews struct {
	gorm.Model
	NewsName    string `json:"news_name" validate:"required"`
	Description string `json:"description" validate:"required"`
}
