package routes

import (
	"html/template"
	"net/http"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	row := db.QueryRow("SELECT id, title, content, created_at FROM posts WHERE id = ?", id)

	var post struct {
		ID        int
		Title     string
		Content   string
		CreatedAt string
	}
	err := row.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/view.html"))
	err = tmpl.Execute(w, post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
