package repository

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/katsukiniwa/practical-go-programming/pkg/entity"
)

type ArticleRepository interface {
	GetArticleList() (article_list []entity.ArticleEntity, err error)
	CreateArticle(article entity.ArticleEntity) (id int, err error)
}

type articleRepository struct{}

func NewArticleRepository() ArticleRepository {
	return &articleRepository{}
}

func (ar *articleRepository) GetArticleList() (article_list []entity.ArticleEntity, err error) {
	article_list = []entity.ArticleEntity{}
	rows, err := Db.Query("SELECT id, title, content FROM article ORDER BY id DESC")
	if err != nil {
		log.Print(err)
		return
	}

	for rows.Next() {
		article := entity.ArticleEntity{}
		err = rows.Scan(&article.Id, &article.Title, &article.Content)
		if err != nil {
			log.Print(err)
			return
		}
		article_list = append(article_list, article)
	}

	return
}

func (tr *articleRepository) CreateArticle(article entity.ArticleEntity) (id int, err error) {
	_, err = Db.Exec("INSERT INTO article (title, content) VALUES (?, ?)", article.Title, article.Content)
	if err != nil {
		log.Print(err)
		return
	}
	err = Db.QueryRow("SELECT id FROM article ORDER BY id DESC LIMIT 1").Scan(&id)

	return
}
