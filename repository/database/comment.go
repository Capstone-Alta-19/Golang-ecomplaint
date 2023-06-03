package database

import (
	"capstone/config"
	"capstone/model"
)

func CreateComment(comment *model.Comment) error {
	err := config.DB.Create(comment).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCommentsByUserID(userID uint) ([]*model.Comment, error) {
	comments := []*model.Comment{}
	err := config.DB.Where("user_id = ?", userID).Find(&comments).Error
	if err != nil {
		return nil, err
	}

	return comments, nil
}
