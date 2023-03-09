package main

import (
	"net/http"
	"time"
)

func getTime(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	var Time = time.Now()
	var response = make(map[string]string)
	response["time"] = Time.Format("2006-01-02T15:04:05Z07:00")
}
