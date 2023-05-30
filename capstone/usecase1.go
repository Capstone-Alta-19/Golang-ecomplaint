package usecase

import (
	"fmt"
	"project_structure/model"
	"project_structure/repository/database"
)

type ReportUsecase interface {
	GetReport(userID uint) (report model.Report, err error)
}

func GetReport(userID uint) (report model.Report, err error) {
	report, err = database.GetReport(userID)
	if err != nil {
		fmt.Println("GetReport: Error getting report from database")
		return
	}
	return
}
