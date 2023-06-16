package usecase

import (
	"capstone/model"
	"capstone/model/payload"
	"capstone/repository/database"
	"errors"
	"fmt"
	"time"
)

func CreateNews(req *payload.CreateNews) (*model.News, error) {
	category, err := database.GetCategoryByID(req.CategoryID)
	if err != nil {
		return nil, errors.New("category not found")
	}
	resp := &model.News{
		NewsName:    req.NewsName,
		Description: req.Description,
		CategoryID:  category.ID,
		Time:        time.Now(),
	}
	err = database.CreateNews(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetListNews() (news []model.News, err error) {
	news, err = database.GetListNews()
	if err != nil {
		fmt.Println("GetListnews : Error getting news from database")
		return
	}
	return
}

func DeleteNews(id uint) (err error) {
	news := model.News{}
	news.ID = id
	err = database.DeleteNews(&news)
	if err != nil {
		fmt.Println("DeleteNews : error deleting news, err: ", err)
		return
	}

	return
}

func UpdateNews(news *model.News) (err error) {
	err = database.UpdateNews(news)
	if err != nil {
		fmt.Println("UpdateNews : Error updating news, err: ", err)
		return
	}

	return
}
