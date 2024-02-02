package model

import (
	"html/template"
	"net/http"
)

// Page struct represents the data to be rendered in the template
type Page struct {
	Title string
}

// RenderTemplate renders an HTML template with the provided data
func RenderTemplate(w http.ResponseWriter, name, tmpl string, data interface{}) {
	t, err := template.New(name + ".html").ParseFiles(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template with data
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
