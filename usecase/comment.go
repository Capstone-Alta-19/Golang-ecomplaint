package usecase

import (
	"capstone/model"
	"capstone/model/payload"
	"capstone/repository/database"
)

func CreateComment(UserID uint, req *payload.CreateCommentRequest) (*model.Comment, error) {

	resp := &model.Comment{
		NewsID:      NewsID,
		Description: req.Description,
	}
	err := database.CreateComment(resp)
	if err != nil {
		return nil, err 
	}
	return resp, nil
}

func GetComments(userID uint) ([]*model.Comment, error) {
	comments, err := database.GetCommentsByUserID(userID)

	if err != nil {
		return nil, err
	}

	return comments, nil
}
