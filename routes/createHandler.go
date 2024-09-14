package routes

import (
	"html/template"
	"net/http"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		content := r.FormValue("content")

		_, err := db.Exec("INSERT INTO posts (title, content) VALUES (?, ?)", title, content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Return the updated list of blog posts
		rows, err := db.Query("SELECT id, title, content, created_at FROM posts ORDER BY created_at DESC")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var posts []struct {
			ID        int
			Title     string
			Content   string
			CreatedAt string
		}

		for rows.Next() {
			var post struct {
				ID        int
				Title     string
				Content   string
				CreatedAt string
			}
			err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			posts = append(posts, post)
		}

		tmpl := template.Must(template.ParseFiles("templates/blog.html"))
		err = tmpl.Execute(w, posts)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/create.html"))
	tmpl.Execute(w, nil)
}
