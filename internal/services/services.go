package services

import (
	"blog-plat/internal/api/v1/schemas"
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
	db.AutoMigrate(&models.Article{})
	return nil
}

func GetUserByID(id float64) (models.User, error) {
	var user models.User
	result := db.Preload("Articles").Where("id = ?", id).First(&user)
	return user, result.Error
}

func GetUserByUsername(username string) (models.User, error) {
	var user models.User
	result := db.Where("username = ?", username).First(&user)
	return user, result.Error
}

func CreateUser(user models.User) (models.User, error) {
	result := db.Create(&user)
	return user, result.Error
}

func GenerateToken(user models.User) (string, error) {
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	godotenv.Load()
	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))
	return token, err
}

func GetArticleList() ([]models.Article, error) {
	var articles []models.Article
	result := db.Preload("User").Find(&articles)
	return articles, result.Error
}

func CreateArticle(article models.Article) (models.Article, error) {
	result := db.Create(&article)
	db.Model(&article).Association("User").Find(&article.User)
	return article, result.Error
}

func GetArticleByID(id int) (models.Article, error) {
	var article models.Article
	result := db.Preload("User").First(&article, id)
	return article, result.Error
}

func UpdateArticleByID(id int, user_id int, updates schemas.UpdateArticleRequest) (models.Article, error) {
	var article models.Article

	if err := db.Where("id = ? AND user_id = ?", id, user_id).First(&article).Error; err != nil {
		return article, err
	}

	if err := db.Model(&article).Updates(updates).Error; err != nil {
		return article, err
	}
	return article, nil
}

func DeleteArticleByID(id int, user_id int) error {
	var article models.Article

	if err := db.Where("id = ? AND user_id = ?", id, user_id).First(&article).Error; err != nil {
		return err
	}

	if err := db.Delete(&article).Error; err != nil {
		return err
	}
	return nil
}
