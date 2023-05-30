package database

import (
	"project_structure/config"
	"project_structure/model"
)

func GetReport(userID uint) (report model.Report, err error) {
	report.UserID = userID
	if err = config.DB.First(&report).Error; err != nil {
		return
	}
	return
}
