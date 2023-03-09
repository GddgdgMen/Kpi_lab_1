package main

import (
	"net/http"
	"strconv"
)

func Start(port int) {
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		errorHandler(err)
	}
}
