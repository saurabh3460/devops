package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	HOST        = "localhost"
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "main"
)

func setupDB() *sql.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dsn)
	checkError(err)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		panic(err)
	}
	return db
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	// sql.Register("postgres", &PSDriver{})
	setupDB()
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Println("hello") })

	log.Fatal(http.ListenAndServe(":8000", router))

}
