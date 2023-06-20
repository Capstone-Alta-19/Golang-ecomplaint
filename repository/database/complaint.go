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

func GetComplaintsByUserID(userID uint, status string) ([]*model.Complaint, error) {
	complaints := []*model.Complaint{}
	DB := config.DB
	DB = DB.Order("created_at desc")
	if status == constant.StatusAll {
		DB = DB.Where("user_id = ?", userID)
	} else {
		DB = DB.Where("user_id = ? AND status = ?", userID, status)
	}
	err := DB.Find(&complaints).Error
	if err != nil {
		return nil, err
	}

	return complaints, nil
}

func GetComplaintByID(id uint) (*model.Complaint, error) {
	var complaint model.Complaint
	err := config.DB.Preload("User").Preload("Category").Preload("Feedback").Preload("Comments.User").Where("id = ?", id).First(&complaint).Error
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

func GetTotalComplaints() (uint, error) {
	var total int64
	err := config.DB.Model(&model.Complaint{}).Where("type = ?", constant.Complaint).Count(&total).Error
	if err != nil {
		return 0, err
	}
	return uint(total), nil
}

func GetTotalAspirations() (uint, error) {
	var total int64
	err := config.DB.Model(&model.Complaint{}).Where("type = ?", constant.Aspiration).Count(&total).Error
	if err != nil {
		return 0, err
	}
	return uint(total), nil
}

func GetAllComplaints(sortby, typeSort, search string, limit, offset int) ([]*model.Complaint, error) {
	complaints := []*model.Complaint{}
	DB := config.DB
	if sortby == constant.Ascending {
		DB = DB.Order("created_at asc")
	}
	if sortby == constant.Descending {
		DB = DB.Order("created_at desc")
	}

	if search != "" {
		DB = DB.Where("title LIKE ?", "%"+search+"%")
	}

	if typeSort == constant.Complaint || typeSort == constant.Aspiration {
		DB = DB.Where("type = ?", typeSort)
	}

	err := DB.Preload("User").Preload("Category").Preload("Feedback").Limit(limit).Offset(offset).Find(&complaints).Error
	if err != nil {
		return nil, err
	}
	return complaints, nil
}

func UpdateComplaint(complaint *model.Complaint) error {
	err := config.DB.Save(complaint).Error
	if err != nil {
		return err
	}
	return nil
}
