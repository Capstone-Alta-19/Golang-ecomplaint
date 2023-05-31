package usecase

import (
	"capstone/model"
	"capstone/repository/database"
	"fmt"
)

func CreateNews(NewsName, Description string) (*model.News, error) {
	resp := &model.News{
		NewsName:    NewsName,
		Description: Description,
	}
	err := database.CreateNews(resp)
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
