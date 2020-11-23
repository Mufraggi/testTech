package app

import (
	"../controller/article_controller"
)

func mapUrls()  {
	router.GET("/articles", article_controller.GetAll)
	router.POST("/article", article_controller.CreateArticle)
	router.GET("/articles/keyWord", article_controller.GetKeysWord)
	router.GET("/articles/dates", article_controller.GetByDate)
}