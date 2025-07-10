package views

import (
	"html/template"
	"log"
)

var Templates *template.Template

func LoadTemplates() {
	var err error
	Templates, err = template.ParseGlob("internal/views/*.gohtml")
	if err != nil {
		log.Fatal("Error loading templates:", err)
	}
}
