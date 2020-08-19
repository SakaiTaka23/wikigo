package models

import (
	"html/template"
	"net/http"
	"regexp"
)

//Page struct
type Page struct {
	Title string
	Body  []byte
}

// Templates var
var Templates = template.Must(template.ParseFiles("views/index.html"))

// ValidPath var
var ValidPath = regexp.MustCompile("^/(edit|save|view|index)/([a-zA-Z0-9]+)$")

// RenderTemplate func
func RenderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := Templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
