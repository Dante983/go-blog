package handlers

import (
	"net/http"

	"github.com/Dante983/go-blog/internal/db"
	"github.com/Dante983/go-blog/internal/views"
)

func AdminNewPostForm(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "New Post - Admin",
	}
	
	AddAuthData(r, data)
	
	err := views.Templates.ExecuteTemplate(w, "admin_new", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AdminCreatePost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	slug := r.FormValue("slug")
	content := r.FormValue("content")

	_, err := db.DB.Exec(`
		INSERT INTO posts (title, slug, content) VALUES (?, ?, ?)`,
		title, slug, content,
	)

	if err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("<p>✅ Post created successfully!</p>"))
}
