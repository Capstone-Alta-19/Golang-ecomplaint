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
	if err = config.DB.Find(&news).Error; err != nil {
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
