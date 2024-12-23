package services

import (
	"blog-plat/internal/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
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

func GetUserByID(id float64)(models.User, error) {
	var user models.User
	result := db.Where("id = ?", id).First(&user)
	return user, result.Error
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


func GenerateToken(user models.User)(string, error){
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	godotenv.Load()
	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))
	return token, err
}