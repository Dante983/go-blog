package middleware

import (
	"net/http"

	"github.com/Dante983/go-blog/internal/auth"
	"github.com/Dante983/go-blog/internal/models"
)

// RequireAuth middleware ensures the user is authenticated
func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !auth.IsAuthenticated(r) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		
		next.ServeHTTP(w, r)
	}
}

// RequireAdmin middleware ensures the user is an admin
func RequireAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := auth.GetUserID(r)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		
		user, err := models.GetUserByID(userID)
		if err != nil || user == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		
		if !user.IsAdmin {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		
		next.ServeHTTP(w, r)
	}
} 