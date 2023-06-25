package database

import (
	"capstone/config"
	"capstone/model"
)

func CreateNews(news *model.News) error {
	if err := config.DB.Create(news).Error; err != nil {
		return err
	}
	return nil
}

func GetListNews() (news []model.News, err error) {
	DB := config.DB
	DB = DB.Order("created_at desc")
	if err = DB.Find(&news).Error; err != nil {
		return
	}
	return
}

func GetFiveNews() (news []model.News, err error) {
	DB := config.DB
	DB = DB.Order("created_at desc").Limit(5)
	if err = DB.Find(&news).Error; err != nil {
		return
	}
	return
}

func DeleteNews(news *model.News) error {
	if err := config.DB.Delete(news).Error; err != nil {
		return err
	}
	return nil
}

func UpdateNews(news *model.News) error {
	if err := config.DB.Updates(news).Error; err != nil {
		return err
	}
	return nil
}

func GetNewsByID(id uint) (*model.News, error) {
	var news model.News
	err := config.DB.Preload("Admin").Preload("Category").Where("id = ?", id).First(&news).Error
	if err != nil {
		return nil, err
	}
	return &news, nil
}
