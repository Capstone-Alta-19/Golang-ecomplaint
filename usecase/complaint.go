package usecase

import (
	"capstone/constant"
	"capstone/model"
	"capstone/model/payload"
	"capstone/repository/database"
)

func CreateComplaint(UserID uint, req *payload.CreateComplaintRequest) (*model.Complaint, error) {

	resp := &model.Complaint{
		UserID:      UserID,
		Description: req.Description,
		Type:        req.Type,
		PhotoURL:    req.PhotoURL,
		VideoURL:    req.VideoURL,
		CategoryID:  req.CategoryID,
		IsPublic:    req.IsPublic,
		Status:      constant.StatusPending,
	}
	err := database.CreateComplaint(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetComplaints(userID uint) ([]*model.Complaint, error) {
	complaints, err := database.GetComplaintsByUserID(userID)

	if err != nil {
		return nil, err
	}

	return complaints, nil
}

func GetComplaintByID(id uint) (*model.Complaint, error) {
	complaint, err := database.GetComplaintByID(id)
	if err != nil {
		return nil, err
	}
	return complaint, nil
}
