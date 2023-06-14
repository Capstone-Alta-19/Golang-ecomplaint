package database

import (
	"capstone/config"
	"capstone/model"
)

func GetPinnedByComplaintIdAndUserId(userID, complaintID uint) (*model.PinnedComplaint, error) {
	var pinned model.PinnedComplaint
	err := config.DB.Where("user_id = ? AND complaint_id = ?", userID, complaintID).First(&pinned).Error
	if err != nil {
		return nil, err
	}
	return &pinned, nil
}

func AddPinnedComplaint(pinned *model.PinnedComplaint) error {
	err := config.DB.Create(pinned).Error
	if err != nil {
		return err
	}
	return nil
}

func DeletePinnedComplaint(pinned *model.PinnedComplaint) error {
	err := config.DB.Delete(pinned).Error
	if err != nil {
		return err
	}
	return nil
}
