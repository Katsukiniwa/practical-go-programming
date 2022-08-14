package router

import (
	"net/http"

	"github.com/katsukiniwa/practical-go-programming/pkg/controller"
)

// 外部パッケージに公開するインターフェース
type ArticleRouter interface {
	HandleArticleRequest(w http.ResponseWriter, r *http.Request)
}

// 非公開のRouter構造体
type articleRouter struct {
	ac controller.ArticleController
}

// Routerのコンストラクタ
// 引数にTaskControllerを受け取りRouter構造体のポインタを返却する
func NewRouter(ac controller.ArticleController) ArticleRouter {
	return &articleRouter{ac}
}

func (ro *articleRouter) HandleArticleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ro.ac.GetArticleList(w, r)
	case "POST":
		ro.ac.PostArticle(w, r)
	default:
		w.WriteHeader(405)
	}
}
