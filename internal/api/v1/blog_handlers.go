package v1

import (
	"blog-plat/internal/api/v1/schemas"
	"blog-plat/internal/models"
	"blog-plat/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ArticleList(c *gin.Context) {
	articles, _ := services.GetArticleList()

	c.JSON(http.StatusOK, gin.H{
		"data": articles,
	})
}

func CreateArticle(c *gin.Context) {
	user, _ := c.Get("currentUser")
	userObj, _ := user.(models.User)

	var req schemas.CreateArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article := models.Article{
        Title: req.Title,
        Body:  req.Body,
		UserID: userObj.ID,
	}
	_, err := services.CreateArticle(article)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func ArticleByID(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	article, err := services.GetArticleByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": article,
	})
}

func UpdateArticleByID(c *gin.Context) {
	user, _ := c.Get("currentUser")
	userObj, _ := user.(models.User)

	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var req schemas.UpdateArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := services.UpdateArticleByID(id, userObj.ID, req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func DeleteArticleByID(c *gin.Context) {
	user, _ := c.Get("currentUser")
	userObj, _ := user.(models.User)

	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	err := services.DeleteArticleByID(id, userObj.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}