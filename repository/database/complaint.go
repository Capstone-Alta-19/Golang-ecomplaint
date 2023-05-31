package database

import (
	"capstone/config"
	"capstone/model"
)

func CreateComplaint(complaint *model.Complaint) error {
	err := config.DB.Create(complaint).Error
	if err != nil {
		return err
	}
	return nil
}

func GetComplaint(complaint *model.Complaint) error {
	err := config.DB.Find(complaint).Error
	if err != nil {
		return err
	}
	return nil
}
