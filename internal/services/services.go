package services

import (
	"blog-plat/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(databaseURL string) error {
	var err error
	db, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return err
	}
	db.AutoMigrate(&models.User{})
	return nil
}

func GetUserByUsername(username string)(models.User, error) {
	var user models.User
	result := db.Where("username = ?", username).First(&user)
	return user, result.Error
}

func CreateUser(user models.User)error{
	result := db.Create(&user)
	return result.Error
}