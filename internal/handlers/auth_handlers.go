package handlers

import (
	"blog-plat/internal/handlers/requests"
	"blog-plat/internal/models"
	"blog-plat/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


func Register(c *gin.Context) {
	var registerInput requests.RegisterInput
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

	
	if err := services.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})

}