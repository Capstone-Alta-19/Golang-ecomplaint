package model

import "gorm.io/gorm"

type Report struct {
	gorm.Model
	Title       string     `json:"title"`
	User        string     `json:"user"`
	Description string     `json:"description"`
	Feedbacks   []Feedback `json:"feedbacks" gorm:"foreignKey:LaporanID"`
}

type Feedback struct {
	gorm.Model
	LaporanID   uint   `json:"-"`
	Description string `json:"description"`
}
