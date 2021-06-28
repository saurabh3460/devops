package router

import (
	"github.com/gorilla/mux"
	"gitlab.com/saurabh3460/go_psql/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/newuser", middleware.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/{id}", middleware.GetUser).Methods("GET")

	return router
}
