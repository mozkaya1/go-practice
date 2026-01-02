package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	log.Println("Starting server on :6060")
	log.Println(http.ListenAndServe("localhost:6060", nil))
}
