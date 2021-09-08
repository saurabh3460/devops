package middleware

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		panic(err)
	}
	return db
}

// func getDbConnection(r *http.Request) *sql.Conn {

// }

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
		log.Fatalf("Unable to connect db at %s %s", r.RequestURI, err)
	}
	defer dbcon.Close()
	var insertID int64
	query := `INSERT INTO Users (name, age, location) VALUES ($1, $2, $3) RETURNING userid`
	if err := db.QueryRow(query, user.Name, user.Age, user.Location).Scan(&insertID); err != nil {
		log.Printf("error found %s", err)
	}
	res := response{
		ID:      insertID,
		Message: "User created successfully",
	}
	json.NewEncoder(w).Encode(res)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var user models.User
	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	dbcon, err := db.Conn(ctx)
	if err != nil {
		log.Fatalf("Unable to connect db at %s %s", r.RequestURI, err)
	}
	defer dbcon.Close()
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Printf("Unable to convert the string into int.  %v", err)
	}
	query := `SELECT * FROM Users WHERE userid = $1`
	ctx, cancel = context.WithTimeout(r.Context(), 1*time.Second)
	defer cancel()
	// time.Sleep(3 * time.Second)
	row := dbcon.QueryRowContext(ctx, query, id)
	if err = row.Scan(&user.ID, &user.Name, &user.Age, &user.Location); err != nil {
		if err == sql.ErrNoRows {
			log.Println("No rows found!")
		} else {
			log.Printf("%v", err)
		}
	}
	json.NewEncoder(w).Encode(user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []models.User
	dbcon := createConnection()
	defer dbcon.Close()
	query := `SELECT * FROM Users`
	rows, err := dbcon.Query(query)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Location)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
}
