package routes

import (
	"html/template"
	"net/http"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	row := db.QueryRow("SELECT id, title, content FROM posts WHERE id = ?", id)

	var post struct {
		ID      string
		Title   string
		Content string
	}
	err := row.Scan(&post.ID, &post.Title, &post.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/view.html"))
	tmpl.Execute(w, post)
}
