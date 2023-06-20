package database

import (
	"capstone/config"
	"capstone/model"
)

func GetCategoryByID(id uint) (*model.Category, error) {
	var category model.Category
	err := config.DB.Where("id = ?", id).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}
