package main

import "log"

func errorHandler(err error) {
	log.Fatal("ListenAndServe: ", err)
}
