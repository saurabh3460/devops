package main

import (
	"log"
	"net/http"
	"os"

	"gitlab.com/saurabh3460/devops/language/go/web/restapi-microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "Product Api: ", log.Lshortfile|log.LstdFlags|log.LUTC)
	handler := handlers.NewHello(l)
	sm := http.NewServeMux()
	sm.Handle("/", handler)
	log.Println("Starting Server")
	err := http.ListenAndServe(":9090", sm)
	log.Fatal(err)
}
