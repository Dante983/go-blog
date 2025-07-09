package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This will list blog posts..."))
}

func SinglePostHandler(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	w.Write([]byte(fmt.Sprintf("Showing blog post: %s", slug)))
}
