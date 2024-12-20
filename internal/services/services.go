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