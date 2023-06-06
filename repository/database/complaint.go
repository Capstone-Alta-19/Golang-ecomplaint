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

func GetComplaintByID(id uint) (*model.Complaint, error) {
	var complaint model.Complaint
	err := config.DB.First(&complaint, id).Error
	if err != nil {
		return nil, err
	}
	return &complaint, nil
}

func GetFeedbacksByComplaintID(complaintID string) ([]*model.Feedback, error) {
	feedbacks := []*model.Feedback{}
	err := config.DB.Where("complaint_id = ?", complaintID).Find(&feedbacks).Error
	if err != nil {
		return nil, err
	}

	return feedbacks, nil
}
