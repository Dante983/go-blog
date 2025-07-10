package handlers

import (
	"net/http"

	"github.com/Dante983/go-blog/internal/views"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Home - Nikola's Blog",
	}
	
	AddAuthData(r, data)
	
	err := views.Templates.ExecuteTemplate(w, "home", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
