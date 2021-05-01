package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	PORT = ":8000"
)

func serveDynamic(w http.ResponseWriter, r *http.Request) {
	response := "The time is now " + time.Now().String()
	fmt.Fprint(w, response)
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/static.html")
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageID := vars["id"]
	fileName := "web/" + pageID + ".html"
	http.ServeFile(w, r, fileName)
}

func accountHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accID := vars["id"]
	fmt.Fprint(w, accID)
}

func main() {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/pages/{id:[0-9]+}", pageHandler)

	rtr.HandleFunc("/static", serveStatic)
	rtr.HandleFunc("/", serveDynamic)

	account := rtr.PathPrefix("/accounts").Subrouter()
	account.HandleFunc("/{id:[0-9]+}", accountHandler)

	http.Handle("/", rtr)
	http.ListenAndServe(PORT, nil)
}
