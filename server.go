package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/katsukiniwa/practical-go-programming/pkg/gateway/middleware"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	var appPort int
	flag.IntVar(&appPort, "port", 9000, "port")
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

	http.ListenAndServe(":9000", nil)
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
