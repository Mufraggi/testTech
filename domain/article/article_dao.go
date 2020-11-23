package article

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type ArticleSchema struct {
	gorm.Model
	Id int64
	Title string
	Content string
	Author string
	Date time.Time
}



func (articleSchema *ArticleSchema) Save() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Create(&articleSchema)
	println("New User Successfully Created")
}

func (articleSchema *ArticleSchema) GetAll() []ArticleSchema {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var articles []ArticleSchema
	db.Find(&articles)
	return articles
}

func (article *ArticleSchema) GetByDate(startDate time.Time, endDate time.Time) []ArticleSchema {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var articles []ArticleSchema
	db.Where("created_at BETWEEN ? AND ?", startDate, endDate).Find(&articles)
	return articles
}

func (articleSchema *ArticleSchema) GetByWord(keyWord string) []ArticleSchema {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var articles []ArticleSchema
	println(keyWord)
	db.Raw("SELECT * FROM article_schemas WHERE instr(content, ?) > 0", keyWord).Find(&articles)
	return articles
}