package controllers

import (
	"html/template"
	"net/http"
)

// IndexHandler func
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/index.html")
	t.Execute(w,r)
}
