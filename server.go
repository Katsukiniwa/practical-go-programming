package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type Comment struct {
	Message  string `validate:"required,min=1,max=140"`
	UserName string `validate:"required,min=1,max=15"`
}

type Book struct {
	Title string `validate:"required"`
	Price int    `validate:"required"`
}

func main() {
	var appPort int
	flag.IntVar(&appPort, "port", 9000, "port")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	fmt.Printf("Server is Ready on :%d\n", appPort)

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
				var out []string
				var ve validator.ValidationErrors
				if errors.As(err, &ve) {
					for _, fe := range ve {
						switch fe.Field() {
						case "Message":
							out = append(out, fmt.Sprintf("Message is 1~140"))
						case "UserName":
							out = append(out, fmt.Sprintf("UserName is 1~15"))
						}
					}
				}
				http.Error(w, fmt.Sprintf(`{"status":"%s"}`, strings.Join(out, ",")), http.StatusBadRequest)
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

	BookImpl()

	http.ListenAndServe(":9000", nil)
}

func BookImpl() {
	s := `{"Title":"Real World HTTP ミニ版", "Price":0}`
	var b Book
	if err := json.Unmarshal([]byte(s), &b); err != nil {
		log.Fatal(err)
	}

	if err := validator.New().Struct(b); err != nil {
		var ve validator.ValidationErrors // validatorの独自型に変換
		if errors.As(err, &ve) {
			for _, fe := range ve {
				fmt.Printf("フィールド %s が %s 違反です(値: %v) \n", fe.Field(), fe.Tag(), fe.Value())
			}
		}
	}
}
