package main

import (
	"context"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 6*time.Second) // it will cancel the operation after one sec
	defer cancel()
	SleepAndTalk(ctx, 5*time.Second, "hello")
}

func SleepAndTalk(ctx context.Context, tm time.Duration, msg string) {
	select {
	case <-time.After(tm):
		fmt.Println(msg)
		return
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		return
	}
}
