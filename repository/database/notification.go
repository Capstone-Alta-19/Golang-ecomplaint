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

func GetNotification() ([]*model.Notification, error) {
	var notifications []*model.Notification
	DB := config.DB
	DB = DB.Order("created_at desc")

	err := DB.Preload("User").Preload("Complaint.Category").Find(&notifications).Error
	if err != nil {
		return nil, err
	}
	return notifications, nil
}
