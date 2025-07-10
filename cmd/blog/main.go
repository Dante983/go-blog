package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	"github.com/Dante983/go-blog/internal/db"
	"github.com/Dante983/go-blog/internal/handlers"
	authMiddleware "github.com/Dante983/go-blog/internal/middleware"
	"github.com/Dante983/go-blog/internal/views"
)

func main() {
	_ = godotenv.Load()
	db.Connect()

	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Static files
	fileServer := http.FileServer(http.Dir("assets"))
	r.Handle("/assets/*", http.StripPrefix("/assets", fileServer))

	// Public routes
	r.Get("/", handlers.HomeHandler)
	r.Get("/posts", handlers.PostsHandler)
	r.Get("/posts/{slug}", handlers.SinglePostHandler)
	r.Get("/load-snippet", handlers.SnippetHandler)

	// Auth routes
	r.Get("/login", handlers.LoginPageHandler)
	r.Post("/login", handlers.LoginHandler)
	r.Get("/logout", handlers.LogoutHandler)

	// Protected admin routes
	r.Get("/admin/posts/new", authMiddleware.RequireAdmin(handlers.AdminNewPostForm))
	r.Post("/admin/posts", authMiddleware.RequireAdmin(handlers.AdminCreatePost))

	views.LoadTemplates()
	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
