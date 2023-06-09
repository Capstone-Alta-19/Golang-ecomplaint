package database

import (
	"capstone/config"
	"capstone/model"
)

func GetLikesCountByComplaintID(complaintID uint) (uint, error) {
	var count int64
	err := config.DB.Model(&model.Like{}).Where("complaint_id = ?", complaintID).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return uint(count), nil
}
