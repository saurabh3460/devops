package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/saurabh3460/devops/language/go/concepts/context/log"
)

func main() {
	http.HandleFunc("/", log.Decorate(handler))
	panic(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, int(12), int64(100))
	log.Println(ctx, "Request at /")
	defer log.Println(ctx, "Request at / finished")
	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "Hello There!")
	case <-ctx.Done():
		err := ctx.Err()
		log.Println(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
