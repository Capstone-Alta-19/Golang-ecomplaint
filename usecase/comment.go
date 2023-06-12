package usecase

import (
	"capstone/model"
	"capstone/model/payload"
	"capstone/repository/database"
)

func CreateCommentByComplaintID(complaintID uint, userID uint, req *payload.CreateCommentRequest) (*model.Comment, error) {
	complaint, err := database.GetComplaintByID(complaintID)
	if err != nil {
		return nil, err
	}

	comment := &model.Comment{
		ComplaintID: complaint.ID,
		UserID:      userID,
		Description: req.Description,
	}

	err = database.CreateComment(comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}
