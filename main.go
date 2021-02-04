package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", httpEcho())

	server := &http.Server{
		Addr:    ":5678",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("ERR server exited with: %s", err)
	}
}

func httpEcho() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://backend-service/")
		if err != nil {
			fmt.Fprintln(w, "hello. Backend returned error: ", err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintln(w, "hello. Backend body read error: ", err)
			return
		}
		fmt.Fprintf(w, "hello. Backend returned: %s", body)
	}
}
