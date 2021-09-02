package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Printf("%s\n", b)
	} else {
		http.Error(w, "Method Not Allowed", 405)
		return
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "OK")
	}
}

func ready(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "OK")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/healthz", health)
	http.HandleFunc("/readyz", ready)
	fmt.Printf("HTTP Post View\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
