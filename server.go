package main

import (
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
	log.Fatal(s.ListenAndServe())
}

func WASMServer(h http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			if strings.HasSuffix(req.URL.Path, ".wasm") {
				resp.Header().Set("content-type", "application/wasm")
			}
		}
		h.ServeHTTP(resp, req)
	})
}
