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

func GetCommentByUserID(userID uint) ([]*model.Comment, error) {
	comments := []*model.Comment{}
	err := config.DB.Preload("User").Where("user_id = ?", userID).Find(&comments).Error
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func DeleteComment(comment *model.Comment) error {
	err := config.DB.Delete(comment).Error
	if err != nil {
		return err
	}
	return nil
}
