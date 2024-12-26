package v1

import (
	"blog-plat/internal/api/v1/schemas"
	"blog-plat/internal/models"
	"blog-plat/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var loginInput schemas.LoginRegisterInput
	if err := c.ShouldBindJSON(&loginInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	userFound, _ = services.GetUserByUsername(loginInput.Username)
	if userFound.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(loginInput.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}

	token, err := services.GenerateToken(userFound)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func Register(c *gin.Context) {
	var registerInput schemas.LoginRegisterInput
	if err := c.ShouldBindJSON(&registerInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	userFound, _ = services.GetUserByUsername(registerInput.Username)
	if userFound.ID != 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "username already used"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(registerInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: registerInput.Username,
		Password: string(passwordHash),
	}

	userObj, err := services.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data": userObj,
	})

}

func GetUserProfile(c *gin.Context) {
	user, _ := c.Get("currentUser")

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
