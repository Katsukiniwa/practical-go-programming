package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type Comment struct {
	Message  string `validate:"required,min=1,max=140"`
	UserName string `validate:"required,min=1,max=15"`
}

func main() {
	var appPort int
	flag.IntVar(&appPort, "port", 9000, "port")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	fmt.Printf("Server is Ready on :%d", appPort)

	var mutex = &sync.RWMutex{}
	comments := make([]Comment, 0, 100)

	http.HandleFunc("/comments", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodGet:
			/*
			 * 読み寄り時に書き込みかあることを考慮しロックする
			 * 本来はDBから読み取る処理を代替
			 */
			mutex.RLock()

			if err := json.NewEncoder(w).Encode(comments); err != nil {
				http.Error(w, fmt.Sprintf(`{"status:"%s"}`, err), http.StatusInternalServerError)
				return
			}

			mutex.RUnlock()

		case http.MethodPost:
			var c Comment
			if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
				http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusInternalServerError)
				return
			}

			/*
			 * バリデーションの追加
			 */
			validate := validator.New()
			if err := validate.Struct(c); err != nil {
				http.Error(w, fmt.Sprintf(`{"status":"%s"`, err), http.StatusBadRequest)
				return
			}

			mutex.Lock()
			// comments := append(comments, c)
			mutex.Unlock()

			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"status": "created"}`))

		default:
			http.Error(w, `{"status":"permits only GET or POST"`, http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	http.ListenAndServe(":9000", nil)
}
