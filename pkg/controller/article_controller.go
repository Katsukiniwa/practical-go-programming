package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "path"
	"strconv"

	"github.com/katsukiniwa/practical-go-programming/pkg/entity"
	"github.com/katsukiniwa/practical-go-programming/pkg/gateway/repository"
)

type ArticleResponse struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ArticleRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ArticleListResponse struct {
	ArticleList []ArticleResponse `json:"article_list"`
}

type ArticleController interface {
	GetArticleList(w http.ResponseWriter, r *http.Request)
	PostArticle(w http.ResponseWriter, r *http.Request)
}

type articleController struct {
	ar repository.ArticleRepository
}

func NewArticleController(ar repository.ArticleRepository) ArticleController {
	return &articleController{ar}
}

func (ac *articleController) GetArticleList(w http.ResponseWriter, r *http.Request) {
	article_list, err := ac.ar.GetArticleList()
	if err != nil {
		w.WriteHeader(500)
		fmt.Print("internal server error")
		return
	}

	var article_response_list []ArticleResponse
	for _, v := range article_list {
		article_response_list = append(article_response_list, ArticleResponse{Id: v.Id, Title: v.Title, Content: v.Content})
	}

	var article_response ArticleListResponse
	article_response.ArticleList = article_response_list

	// JSONに変換
	output, _ := json.MarshalIndent(article_response.ArticleList, "", "\t\t")

	// JSONを返却
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func (ac *articleController) PostArticle(w http.ResponseWriter, r *http.Request) {
	// Request BodyのJSONをDTOにマッピング
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var article_request ArticleRequest
	json.Unmarshal(body, &article_request)

	article := entity.ArticleEntity{Title: article_request.Title, Content: article_request.Content}

	id, err := ac.ar.CreateArticle(article)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	// LocationにリソースのPATHを設定し、ステータスコード２０１を返却
	w.Header().Set("Location", r.Host+r.URL.Path+strconv.Itoa(id))
	w.WriteHeader(201)
}
