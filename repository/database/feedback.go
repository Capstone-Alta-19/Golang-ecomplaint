package database

import (
	"capstone/config"
	"capstone/model"

	"gorm.io/gorm"
)

func GetFeedbackByComplaintID(complaintID uint) (*model.Feedback, error) {
	var feedback model.Feedback
	err := config.DB.Where("complaint_id = ?", complaintID).First(&feedback).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			emptyFeedback := model.Feedback{}
			return &emptyFeedback, nil
		}
		return nil, err
	}
	return &feedback, nil
}
