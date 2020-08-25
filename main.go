package main

import (
	"log"
	"net/http"

	"github.com/SakaiTaka23/wikigo/controllers"
)

func main() {
	http.HandleFunc("/index", controllers.IndexHandler)
	http.HandleFunc("/edit/", controllers.EditHandler)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
