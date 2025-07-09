package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/Dante983/go-blog/internal/db"
)

func main() {
	_ = godotenv.Load()
	db.Connect()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Blog Home - DB Connected!"))
	})

	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", mux)
}
