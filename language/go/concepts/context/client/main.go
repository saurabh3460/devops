package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// func main() {
// 	res, err := http.Get("http://localhost:8000")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// if res.StatusCode != http.StatusOK {
// 	log.Fatal(res.Status)
// }
// defer res.Body.Close()
// io.Copy(os.Stdout, res.Body)
// }

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8000", nil)

	if err != nil {
		log.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		log.Fatalf("status code: %s", res.Status)
	}

	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)

}
