package auth

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

func init() {
	// Use environment variable for session key, or default for development
	sessionKey := os.Getenv("SESSION_KEY")
	if sessionKey == "" {
		// CHANGE THIS IN PRODUCTION!
		sessionKey = "your-32-byte-long-secret-key-here!!"
	}
	
	store = sessions.NewCookieStore([]byte(sessionKey))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 7 days
		HttpOnly: true,
		Secure:   false, // Set to true in production with HTTPS
		SameSite: http.SameSiteLaxMode,
	}
}

// GetSession retrieves the session
func GetSession(r *http.Request) (*sessions.Session, error) {
	return store.Get(r, "blog-session")
}

// SetUserID sets the user ID in the session
func SetUserID(w http.ResponseWriter, r *http.Request, userID int) error {
	session, err := GetSession(r)
	if err != nil {
		return err
	}
	
	session.Values["user_id"] = userID
	return session.Save(r, w)
}

// GetUserID retrieves the user ID from the session
func GetUserID(r *http.Request) (int, bool) {
	session, err := GetSession(r)
	if err != nil {
		return 0, false
	}
	
	userID, ok := session.Values["user_id"].(int)
	return userID, ok
}

// ClearSession clears the session (logout)
func ClearSession(w http.ResponseWriter, r *http.Request) error {
	session, err := GetSession(r)
	if err != nil {
		return err
	}
	
	session.Options.MaxAge = -1
	return session.Save(r, w)
}

// IsAuthenticated checks if the user is logged in
func IsAuthenticated(r *http.Request) bool {
	_, ok := GetUserID(r)
	return ok
} 