package log

import (
	"context"
	"log"
	"math/rand"
	"net/http"
)

type key int

const ReqIDKey = key(12)

func Println(ctx context.Context, msg string) {
	id, ok := ctx.Value(ReqIDKey).(int64)
	if !ok {
		log.Println("could not find request ID")
		return
	}
	log.Printf("[%d] %s", id, msg)
}

func Decorate(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := rand.Int63()
		ctx = context.WithValue(ctx, ReqIDKey, id)
		f(w, r.WithContext(ctx))
	}
}
