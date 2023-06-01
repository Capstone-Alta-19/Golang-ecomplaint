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

func GetComplaintsByUserID(userID uint) ([]*model.Complaint, error) {
	complaints := []*model.Complaint{}
	err := config.DB.Where("user_id = ?", userID).Find(&complaints).Error
	if err != nil {
		return nil, err
	}

	return complaints, nil
}
