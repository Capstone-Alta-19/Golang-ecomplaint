package usecase

import (
	"capstone/model"
	"capstone/model/payload"
	"capstone/repository/database"
	"capstone/utils"
	"errors"
	"fmt"
	"time"
)

func CreateNews(req *payload.CreateNews, adminID uint) (*model.News, error) {
	category, err := database.GetCategoryByID(req.CategoryID)
	if err != nil {
		return nil, errors.New("category not found")
	}
	resp := &model.News{
		NewsName:    req.NewsName,
		PhotoURL:    req.PhotoURL,
		Description: req.Description,
		CategoryID:  category.ID,
		AdminID:     adminID,
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

func GetNewsByID(newsID uint) (*payload.GetNewsByIDResponse, error) {
	news, err := database.GetNewsByID(newsID)
	if err != nil {
		return nil, errors.New("news not found")
	}

	fiveNews, err := database.GetFiveNews()
	if err != nil {
		return nil, err
	}

	newsList := []payload.NewsList{}
	for _, v := range fiveNews {
		if v.ID != news.ID {
			newsList = append(newsList, payload.NewsList{
				ID:       v.ID,
				NewsName: v.NewsName,
				PhotoURL: v.PhotoURL,
			})
		}
	}

	resp := &payload.GetNewsByIDResponse{
		ID:          news.ID,
		NewsName:    news.NewsName,
		PhotoURL:    news.PhotoURL,
		Description: news.Description,
		Admin:       news.Admin.Name,
		Category:    news.Category.Name,
		CreatedAt:   utils.ConvertDateToIndonesia(news.Time.Format("Monday, 02 Januari 2006")),
		NewsList:    newsList,
	}

	return resp, nil
}
