package database

import (
	"capstone/config"
	"capstone/model"
)

func CreateComment(comment *model.Comment) error {
	err := config.DB.Create(comment).Error
	if err != nil {
		return err
	}
	return nil
}
