package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	var appPort int
	flag.IntVar(&appPort, "port", 9000, "port")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	fmt.Printf("Server is Ready on :%d\n", appPort)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	http.HandleFunc("/params", handleParams)

	http.ListenAndServe(":9000", nil)
}

func handleParams(w http.ResponseWriter, r *http.Request) {

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
