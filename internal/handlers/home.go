package handlers

import (
	"net/http"

	"github.com/Dante983/go-blog/internal/views"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := views.Templates.ExecuteTemplate(w, "layout.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
