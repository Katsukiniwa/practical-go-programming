package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/katsukiniwa/practical-go-programming/pkg/controller"
	"github.com/katsukiniwa/practical-go-programming/pkg/gateway/middleware"
	"github.com/katsukiniwa/practical-go-programming/pkg/gateway/repository"
	"github.com/katsukiniwa/practical-go-programming/pkg/gateway/router"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	var appPort int
	flag.IntVar(&appPort, "port", 9001, "port")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	fmt.Printf("Server is Ready on :%d\n", appPort)

	yes := 0
	no := 0

	r := chi.NewRouter()

	/*
	 * answer handler
	 */
	r.Post("/poll/{answer}", func(w http.ResponseWriter, r *http.Request) {
		if chi.URLParam(r, "answer") == "y" {
			yes++
		} else {
			no++
		}
	})

	r.Get("/result", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "賛成: %d, 反対: %d", yes, no)
	})
	r.Handle("/asset/*", http.StripPrefix("/asset/", http.FileServer(http.Dir("."))))

	http.HandleFunc("/", RootHandler)

	http.HandleFunc("/params", ParamsHandler)

	http.Handle("/health", middleware.MiddlewareLogging(http.HandlerFunc(middleware.HealthCheck)))

	http.Handle("/health/new", middleware.WrapHandlerWithLogging(http.HandlerFunc(middleware.HealthCheck)))

	http.Handle("/error", middleware.Recovery(http.HandlerFunc(middleware.PanicHealthCheck)))

	var tr = repository.NewArticleRepository()
	var tc = controller.NewArticleController(tr)
	var ro = router.NewRouter(tc)
	http.HandleFunc("/articles/", ro.HandleArticleRequest)

	UserRequest()

	http.ListenAndServe(":9001", nil)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}

func ParamsHandler(w http.ResponseWriter, r *http.Request) {

	// クエリパラメータ取得してみる
	fmt.Fprintf(w, "クエリ：%s\n", r.URL.RawQuery)

	// Bodyデータを扱う場合には、事前にパースを行う
	r.ParseForm()

	// Formデータを取得.
	form := r.PostForm
	fmt.Fprintf(w, "フォーム：\n%v\n", form)

	// または、クエリパラメータも含めて全部.
	params := r.Form
	fmt.Fprintf(w, "フォーム2: \n%v\n", params)
}

type User struct {
	Name    string
	Address string
}

func UserRequest() {
	u := User{
		Name:    "オライリー",
		Address: "東京都新宿区四谷坂町",
	}

	payload, err := json.Marshal(u)
	if err != nil {
		fmt.Print(err)
	}

	resp, err := http.Post("http://example.com", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Print(err)
	}

	getResp, err := http.Get("https://next-rails-playground.herokuapp.com/categories")
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer getResp.Body.Close()
	byteArray, _ := ioutil.ReadAll(getResp.Body)
	stringByteArray := string(byteArray)
	fmt.Println(stringByteArray)

	defer resp.Body.Close()
}
