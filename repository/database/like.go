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

func GetLikeByComplaintIdAndUserId(userID uint, complaintID uint) (*model.Like, error) {
	var like model.Like
	err := config.DB.Where("user_id = ? AND complaint_id = ?", userID, complaintID).First(&like).Error
	if err != nil {
		return nil, err
	}
	return &like, nil
}

func AddLikeByComplaintID(like *model.Like) error {
	err := config.DB.Create(like).Error
	if err != nil {
		return err
	}
	return nil
}
