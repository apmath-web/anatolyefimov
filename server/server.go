package main

import (
	"fmt"
	"net/http"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "main page")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
