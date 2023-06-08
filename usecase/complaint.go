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
		Type:        req.Type,
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

func GetComplaintByID(id uint) (*model.Complaint, error) {
	complaint, err := database.GetComplaintByID(id)
	if err != nil {
		return nil, err
	}
	return complaint, nil
}

func GetFeedback(complaintID string) ([]*model.Feedback, error) {
	feedbacks, err := database.GetFeedbacksByComplaintID(complaintID)
	if err != nil {
		return nil, err
	}

	return feedbacks, nil
}

func GetComplaintsByCategoryId(categoryID uint, sortParam string) ([]*model.Complaint, error) {
	complaints, err := database.GetComplaintsByCategoryAndSort(categoryID, sortParam)
	if err != nil {
		return nil, err
	}

	return complaints, nil
}

func GetLikes(sortBy string, query string) ([]*model.Like, error) {
	likes, err := database.GetLikes(sortBy, query)
	if err != nil {
		return nil, err
	}

	return likes, nil
}
