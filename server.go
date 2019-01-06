package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	fs := http.FileServer(http.Dir("./html"))
	s := &http.Server{
		Addr:         ":8080",
		Handler:      WASMServer(fs),
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}
	fmt.Println("open your browser on http://127.0.0.1:8080")
	log.Fatal(s.ListenAndServe())
}

func WASMServer(h http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			if strings.HasSuffix(req.URL.Path, ".wasm") {
				resp.Header().Set("content-type", "application/wasm")
			}
		}

		resp.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		resp.Header().Set("Pragma", "no-cache")
		resp.Header().Set("Expires", "0")

		h.ServeHTTP(resp, req)
	})
}
