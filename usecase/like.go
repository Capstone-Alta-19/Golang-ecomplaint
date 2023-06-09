package usecase

import (
	"capstone/model"
	"capstone/repository/database"
	"errors"

	"gorm.io/gorm"
)

func LikeByComplaintID(userID uint, complaintID uint64) error {
	complaint, err := database.GetComplaintByID(uint(complaintID))
	if err != nil {
		return errors.New("complaint not found")
	}
	liked, err := database.GetLikeByComplaintIdAndUserId(userID, complaint.ID)
	if err != nil && err == errors.New("record not found") {
		return err
	}
	if liked != nil {
		return errors.New("you have liked this complaint")
	}

	like := model.Like{
		UserID:      userID,
		ComplaintID: complaint.ID,
	}

	err = database.AddLikeByComplaintID(&like)
	if err != nil {
		return err
	}

	return nil
}

func UnLikeByComplaintID(userID uint, complaintID uint64) error {
	complaint, err := database.GetComplaintByID(uint(complaintID))
	if err != nil {
		return errors.New("complaint not found")
	}

	liked, err := database.GetLikeByComplaintIdAndUserId(userID, complaint.ID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if liked == nil {
		return errors.New("you have not liked this complaint")
	}

	err = database.UnLikeByComplaintIdAndUserId(userID, complaint.ID)
	if err != nil {
		return err
	}

	return nil
}
