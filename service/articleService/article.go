package articleService

import (
	"../../domain/article"
	"../../utils/date_utils"
)
func CreateArticle(articleData article.Article) (*article.Article) {
	var newArticle article.ArticleSchema
	newArticle.Content = articleData.Content
	newArticle.Date = date_utils.GetNowString()
	newArticle.Author = articleData.Author
	newArticle.Title = articleData.Titre
	newArticle.Save()
	return &articleData
}

func GetAll()  []article.Article {
	var newArticle article.ArticleSchema
	var res []article.Article
	result := newArticle.GetAll()
	for i := 0; i < len(result); i++ {
		articleTmp := result[i]
		tmp := article.Article{
			Titre: articleTmp.Title,
			Content: articleTmp.Content,
			Author: articleTmp.Author,
			Date: articleTmp.Date.String(),
		}
		res = append(res, tmp)
	}
	return res
}

func GetByKeyword(getKeywordDto article.GetKeywordDto) []article.Article {
	keyword := getKeywordDto.Word
	var newArticle article.ArticleSchema
	var res []article.Article
	result := newArticle.GetByWord(keyword)
	println(len(result))
	for i := 0; i < len(result); i++ {
		articleTmp := result[i]
		tmp := article.Article{
			Titre: articleTmp.Title,
			Content: articleTmp.Content,
			Author: articleTmp.Author,
			Date: articleTmp.Date.String(),
		}
		res = append(res, tmp)
	}
	return res
}

func GetByDate(getByDate article.GetByDate) []article.Article {
	start := date_utils.FormatDateToTime(getByDate.Start)
	end :=  date_utils.FormatDateToTime(getByDate.End)

	var newArticle article.ArticleSchema
	var res []article.Article
	result := newArticle.GetByDate(start,end)
	for i := 0; i < len(result); i++ {
		articleTmp := result[i]
		tmp := article.Article{
			Titre: articleTmp.Title,
			Content: articleTmp.Content,
			Author: articleTmp.Author,
			Date: articleTmp.Date.String(),
		}
		res = append(res, tmp)
	}
	return res
}