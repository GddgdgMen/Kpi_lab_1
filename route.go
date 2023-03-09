package main

import "net/http"

func main() {
	http.HandleFunc("/time", getTime)

	Start(8795)
}
