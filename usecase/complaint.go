package usecase

import (
	"capstone/model"
	"capstone/model/payload"
	"capstone/repository/database"
)

func CreateComplaint(UserID uint, req *payload.CreateComplaintRequest) (*model.Complaint, error) {

	resp := &model.Complaint{
		UserID:      UserID,
		Description: req.Description,
		Type:        "Complaint",
		CategoryID:  req.CategoryID,
		IsPublic:    req.IsPublic,
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
