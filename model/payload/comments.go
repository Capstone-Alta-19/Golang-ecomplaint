package payload

import (
	"gorm.io/gorm"
)

type CreateComment struct {
	gorm.Model
	NewsID      string `json:"news_id" validate:"required"`
	Description string `json:"description" validate:"required"`
}
