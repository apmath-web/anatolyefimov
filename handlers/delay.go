package handlers

import (
	"fmt"
	"net/http"
	"time"
	"strconv"
)

func Wait(w http.ResponseWriter, r *http.Request) {
	delay := r.URL.Query()["delay"][0]
	duration, err := strconv.Atoi(delay)
	if  err == nil {
		time.Sleep(time.Duration(duration)* time.Millisecond)
	}
	fmt.Fprintf(w, "Sleep for %sms completed", delay)
}
