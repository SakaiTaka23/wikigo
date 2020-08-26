package main

import (
	"log"
	"net/http"

	"github.com/SakaiTaka23/wikigo/controllers"
)

func main() {
	http.HandleFunc("/index", controllers.IndexHandler)
	http.HandleFunc("/edit/", controllers.EditHandler)
	http.HandleFunc("/show/", controllers.ShowHandler)
	http.HandleFunc("/update/", controllers.UpdateHandler)
	http.HandleFunc("/create", controllers.CreateHandler)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
