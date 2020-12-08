package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/ping", handler2)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey how's it going?")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Insert meme here")
}
