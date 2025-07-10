package handlers

import (
	"net/http"

	"github.com/Dante983/go-blog/internal/auth"
	"github.com/Dante983/go-blog/internal/models"
	"github.com/Dante983/go-blog/internal/views"
)

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	// If already logged in, redirect to admin
	if auth.IsAuthenticated(r) {
		http.Redirect(w, r, "/admin/posts/new", http.StatusSeeOther)
		return
	}

	data := map[string]interface{}{
		"Title": "Login - Nikola's Blog",
	}

	err := views.Templates.ExecuteTemplate(w, "login", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	// Get user by username
	user, err := models.GetUserByUsername(username)
	if err != nil {
		w.Write([]byte(`<div style="color: #ff0000; border: 1px solid #ff0000; padding: 10px;">
			> ERROR: System failure. Try again.
		</div>`))
		return
	}

	if user == nil || !models.CheckPasswordHash(password, user.PasswordHash) {
		w.Write([]byte(`<div style="color: #ff0000; border: 1px solid #ff0000; padding: 10px;">
			> ACCESS DENIED: Invalid credentials
		</div>`))
		return
	}

	// Set session
	if err := auth.SetUserID(w, r, user.ID); err != nil {
		w.Write([]byte(`<div style="color: #ff0000; border: 1px solid #ff0000; padding: 10px;">
			> ERROR: Session initialization failed
		</div>`))
		return
	}

	// Success - redirect with HTMX
	w.Header().Set("HX-Redirect", "/admin/posts/new")
	w.Write([]byte(`<div style="color: #00ff00; border: 1px solid #00ff00; padding: 10px;">
		> ACCESS GRANTED
		> Redirecting to admin panel...
	</div>`))
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if err := auth.ClearSession(w, r); err != nil {
		http.Error(w, "Failed to logout", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
