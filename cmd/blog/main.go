package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	"github.com/Dante983/go-blog/internal/db"
	"github.com/Dante983/go-blog/internal/handlers"
)

func main() {
	_ = godotenv.Load()
	db.Connect()

	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Routes
	r.Get("/", handlers.HomeHandler)
	r.Get("/posts", handlers.PostsHandler)
	r.Get("/posts/{slug}", handlers.SinglePostHandler)

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
