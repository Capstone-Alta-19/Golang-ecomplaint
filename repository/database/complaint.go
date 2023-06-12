package database

import (
	"capstone/config"
	"capstone/constant"
	"capstone/model"
)

func CreateComplaint(complaint *model.Complaint) error {
	err := config.DB.Create(complaint).Error
	if err != nil {
		return err
	}
	return nil
}

func GetComplaintsByCategoryID(categoryID uint, sort string) ([]*model.Complaint, error) {
	complaints := []*model.Complaint{}
	DB := config.DB
	if sort == constant.Ascending {
		DB = DB.Order("created_at asc")
	} else if sort == constant.Descending {
		DB = DB.Order("created_at desc")
	}

	err := DB.Preload("User").Preload("Category").Preload("Feedback").Where("category_id = ?", categoryID).Find(&complaints).Error
	if err != nil {
		return nil, err
	}

	return complaints, nil
}

func GetComplaintsByUserID(userID uint) ([]*model.Complaint, error) {
	complaints := []*model.Complaint{}
	err := config.DB.Where("user_id = ?", userID).Find(&complaints).Error
	if err != nil {
		return nil, err
	}

	return complaints, nil
}

func GetComplaintByID(id uint) (*model.Complaint, error) {
	var complaint model.Complaint
	err := config.DB.First(&complaint, id).Error
	if err != nil {
		return nil, err
	}
	return &complaint, nil
}

func DeleteComplaint(complaint *model.Complaint) error {
	err := config.DB.Delete(complaint).Error
	if err != nil {
		return err
	}
	return nil
}
