package usecase

import (
	"capstone/model"
	"capstone/repository/database"
	"fmt"
)

func CreateComment(NewsID, Description string) (*model.Comment, error) {
	resp := &model.Comment{
		NewsID:      NewsID,
		Description: Description,
	}
	err := database.CreateComment(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetComment(id uint) (comment model.Comment, err error) {
	comment, err = database.GetComment(id)
	if err != nil {
		fmt.Println("GetComment : Error getting comment from database")
		return
	}
	return
}

func GetListComment() (comment []model.Comment, err error) {
	comment, err = database.GetListComment()
	if err != nil {
		fmt.Println("GetListcomment : Error getting comment from database")
		return
	}
	return
}

func DeleteComment(id uint) (err error) {
	comment := model.Comment{}
	comment.ID = id
	err = database.DeleteComment(&comment)
	if err != nil {
		fmt.Println("DeleteComment : error deleting comment, err: ", err)
		return
	}

	return
}

func UpdateComment(comment *model.Comment) (err error) {
	err = database.UpdateComment(comment)
	if err != nil {
		fmt.Println("UpdateComment : Error updating comment, err: ", err)
		return
	}

	return
}
