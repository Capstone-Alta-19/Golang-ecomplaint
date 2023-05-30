package config

import (
	"Golang-ecomplaint/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "",
		"DB_Port":     "3306",
		"DB_Host":     "localhost",
		"DB_Name":     "e-complaint",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"])

	var e error
	DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
	InitCategory()
	return DB
}

func InitMigrate() {
	DB.AutoMigrate(&model.Category{})
}

func InitCategory() {
	categories := []model.Category{
		{Name: "Dosen dan Staff Akademik"},
		{Name: "Sarana dan Prasarana"},
		{Name: "Sistem Perkuliahan"},
		{Name: "Organisasi Mahasiswa"},
		{Name: "Sesama Mahasiswa"},
		{Name: "Lainnya"},
	}
	for _, category := range categories {
		var exist model.Category
		err := DB.Where("name = ?", category.Name).First(&exist).Error
		if err != nil {
			DB.Create(&category)
		}
	}
}
