package config

import (
	"capstone/constant"
	"capstone/model"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() *gorm.DB {

	config := Config{
		DB_Username: os.Getenv("DB_USER"),
		DB_Password: os.Getenv("DB_PASSWORD"),
		DB_Port:     os.Getenv("DB_PORT"),
		DB_Host:     os.Getenv("DB_HOST"),
		DB_Name:     os.Getenv("DB_NAME"),
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	InitialMigration()
	InitCategory()
	AdminSeed()
	return DB
}

func InitialMigration() {
	DB.AutoMigrate(
		&model.User{},
		&model.Admin{},
		&model.Complaint{},
		&model.News{},
		&model.Category{},
		&model.Comment{},
		&model.Like{},
		&model.Feedback{},
		&model.PinnedComplaint{},
		&model.Notification{},
	)
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

func AdminSeed() {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("ADMIN_PASSWORD")), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	admin := model.Admin{
		Name:     "Super Admin",
		Role:     constant.SuperAdmin,
		Username: "admin",
		Password: string(passwordHash),
	}
	var exist model.Admin
	err = DB.Where("username = ?", admin.Username).First(&exist).Error
	if err != nil {
		DB.Create(&admin)
	}
}
