package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func Date(w http.ResponseWriter, r* http.Request) {
	fmt.Fprintf(w, "Current date and time: %s", time.Now())
} 
