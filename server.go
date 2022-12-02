package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
	"github.com/katsukiniwa/practical-go-programming/ent"
	"github.com/katsukiniwa/practical-go-programming/ent/migrate"
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

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	client, err := ent.Open(dialect.MySQL, dsn)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	if err := client.Schema.Create(ctx, migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}
	log.Print("ent sample done.")

	cmp, err := client.Debug().Company.
		Create().
		SetName("companyA").
		Save(ctx)
	if err != nil {
		log.Fatalf("failed create company: %v", err)
	}
	log.Printf("cmp: %+v", cmp)
	usr, err := client.Debug().User.
		Create().
		SetFirstName("first name").
		SetLastName("last name").
		SetAge(20).
		SetEmail("example@example.co.jp").
		SetCompany(cmp).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed create user: %v", err)
	}
	log.Printf("user: %+v", usr)

	log.Print("ent sample done.")
	http.ListenAndServe(":9001", nil)
}
