package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/SakaiTaka23/wikigo/models"
)

// IndexHandler func
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	res := models.GetIndex()
	t, _ := template.ParseFiles("views/index.html")
	t.Execute(w, res)
}

// ShowHandler func
func ShowHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/show/"):]
	fmt.Println("title:" + title)
	res := models.GetInfo(title)
	t, _ := template.ParseFiles("views/show.html")
	t.Execute(w, res)
}

// EditHandler func
func EditHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	a := models.GetInfo(title)
	t, _ := template.ParseFiles("views/edit.html")
	t.Execute(w, a)
}

// CreateHandler func
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/create.html")
	t.Execute(w, r)
}

// UpdateHandler func
func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	if title == "" {
		title = r.URL.Path[len("/update/"):]
	}
	body := r.FormValue("body")
	author := r.FormValue("author")
	p := models.Article{Title: title, Body: body, Author: author, UpdatedAt: time.Now()}
	fmt.Println("p:")
	fmt.Println(p)
	p.SaveInfo()
	fmt.Println("title:" + title)
	http.Redirect(w, r, "/show/"+title, http.StatusFound)
}
