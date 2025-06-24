package main

import (
	"net/http"
	"log"
)

func main() {
	r := InitRouter()
	log.Fatal(http.ListenAndServe(":1234", r))
}
