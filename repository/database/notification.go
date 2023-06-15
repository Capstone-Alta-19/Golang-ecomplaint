package database

import (
	"capstone/config"
	"capstone/model"
)

func CreateNotification(notification *model.Notification) error {
	err := config.DB.Create(notification).Error
	if err != nil {
		return err
	}
	return nil
}
