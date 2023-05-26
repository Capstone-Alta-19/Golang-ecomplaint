package config

import (
	"capstone/model"
	"fmt"
	"os"

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
		DB_Username: os.Getenv("DB_Username"),
		DB_Password: os.Getenv("DB_Password"),
		DB_Port:     os.Getenv("DB_Port"),
		DB_Host:     os.Getenv("DB_Host"),
		DB_Name:     os.Getenv("DB_Name"),
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
	return DB
}

func InitialMigration() {
	DB.AutoMigrate(
		&model.User{},
		&model.Admin{},
		&model.Complaint{},
	)
}
