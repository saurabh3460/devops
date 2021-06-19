package middleware

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
	"gitlab.com/saurabh3460/go_psql/models"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

var db = createConnection()

func createConnection() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/go_psql?sslmode=disable")
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	return db
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	dbcon, err := db.Conn(ctx)
	if err != nil {
		log.Fatalf("Unable to connect db in CreateUser %s", err)
	}
	defer dbcon.Close()
	var insertID int64
	sqlQuery := `INSERT INTO Users (name, age, location) VALUES ($1, $2, $3) RETURNING userid`
	if err := db.QueryRow(sqlQuery, user.Name, user.Age, user.Location).Scan(&insertID); err != nil {
		log.Printf("error found %s", err)
	}
	res := response{
		ID:      insertID,
		Message: "User created successfully",
	}
	json.NewEncoder(w).Encode(res)
}
