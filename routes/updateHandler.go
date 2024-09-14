package routes

import (
	"html/template"
	"net/http"
)

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		title := r.FormValue("title")
		content := r.FormValue("content")

		_, err := db.Exec("UPDATE posts SET title = ?, content = ? WHERE id = ?", title, content, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/blog", http.StatusSeeOther)
		return
	}

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

	tmpl := template.Must(template.ParseFiles("templates/update.html"))
	tmpl.Execute(w, post)
}
