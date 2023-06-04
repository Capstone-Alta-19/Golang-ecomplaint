package database

import (
	"capstone/config"
	"capstone/model"
)

func CreateComment(comment *model.Comment) error {
	if err := config.DB.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

func GetComments() (interface{}, error) {
	var comments []model.Comment

	if err := config.DB.Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func GetComment(id uint) (comment model.Comment, err error) {
	comment.ID = id
	if err = config.DB.First(&comment).Error; err != nil {
		return
	}
	return
}

func GetListComment() (comment []model.Comment, err error) {
	if err = config.DB.Find(&comment).Error; err != nil {
		return
	}
	return
}

func DeleteComment(comment *model.Comment) error {
	if err := config.DB.Delete(comment).Error; err != nil {
		return err
	}
	return nil
}

func UpdateComment(comment *model.Comment) error {
	if err := config.DB.Updates(comment).Error; err != nil {
		return err
	}
	return nil
}
