package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "v3 %d", n)
	})

	(&http.Server{
		Addr:    ":3000",
		Handler: mux,
	}).ListenAndServe()
}
