package controllers

import (
	"html/template"
	"net/http"
	"github.com/SakaiTaka23/wikigo/models"
)

// IndexHandler func
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	res := models.GetIndex()
	t, _ := template.ParseFiles("views/index.html")
	t.Execute(w, res)
}

// EditHandler func
func EditHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	a := models.Find(title)
	t, _ := template.ParseFiles("views/edit.html")
	t.Execute(w, a)
}

// SaveHandler func
func SaveHandler(w http.ResponseWriter,r *http.Request){
	title := r.FormValue("title")
	body := r.FormValue("body")
	author := r.FormValue("author")
	p := models.Article{Title: title,Body:body,Author:author}
	p.Save()
}
