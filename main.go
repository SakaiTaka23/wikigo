package main

import (
	"log"
	"net/http"
	"controllers"
)

func main() {
	http.HandleFunc("/index", IndexHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
