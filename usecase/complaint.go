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
