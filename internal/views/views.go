package views

import (
	"html/template"
	"log"
)

var Templates *template.Template

func LoadTemplates() {
	var err error
	Templates, err = template.ParseFiles(
			"internal/views/layout.gohtml",
			"internal/views/home.gohtml",
			"internal/views/admin_new.gohtml",
	)
	if err != nil {
		log.Fatal("Error loading templates:", err)
	}
}
