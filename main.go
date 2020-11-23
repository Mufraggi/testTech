
package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"./app"
	"./domain/article"
)

func main() {
	InitialMigration()
	app.StartApplication()
}



func InitialMigration() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(article.ArticleSchema{})
}

