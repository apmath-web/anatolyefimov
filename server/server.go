package main

import (
	"net/http"
	"log"
	"github.com/apmath-web/anatolyefimov-server/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Root)
	http.HandleFunc("/date", handlers.Date)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
