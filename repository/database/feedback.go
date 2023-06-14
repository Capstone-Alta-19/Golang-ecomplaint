package database

import (
	"capstone/config"
	"capstone/model"
)

func CreateFeedback(feedback *model.Feedback) error {
	err := config.DB.Create(feedback).Error
	if err != nil {
		return err
	}
	return nil
}
