package article_controller

import (
	"../../domain/article"
	"../../service/articleService"
	"github.com/gin-gonic/gin"
	"net/http"
)
func CreateArticle(c *gin.Context)  {
	var article article.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, "bad_request")
		return
	}
	result := articleService.CreateArticle(article)
	c.JSON(http.StatusCreated, result)
	return
}

func GetAll(c *gin.Context) {
	result := articleService.GetAll()
	c.JSON(http.StatusCreated, result)
	return
}

func GetKeysWord(c *gin.Context) {
	var getKeywordDto article.GetKeywordDto
	if err := c.ShouldBindJSON(&getKeywordDto); err != nil {
		c.JSON(http.StatusBadRequest, "bad_request")
		return
	}
	result := articleService.GetByKeyword(getKeywordDto)
	c.JSON(http.StatusCreated, result)
}

func GetByDate(c *gin.Context) {
	var getByDate article.GetByDate
	if err := c.ShouldBindJSON(&getByDate); err != nil {
		c.JSON(http.StatusBadRequest, "bad_request")
		return
	}
	result := articleService.GetByDate(getByDate)
	c.JSON(http.StatusCreated, result)
}