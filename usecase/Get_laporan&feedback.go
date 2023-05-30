package usecase

import (
	"capstone/model"
	"capstone/model/payload"
	"capstone/repository/database"
)

func GetComplaint(UserID uint, req *payload.GetComplaintRequest) (*model.Complaint, error) {
	resp := &model.Complaint{
		UserID:      UserID,
		Description: req.Description,
		Type:        "Complaint",
		CategoryID:  req.CategoryID,
		Feedback:    req.Feedback,
	}

	err := database.GetComplaint(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
