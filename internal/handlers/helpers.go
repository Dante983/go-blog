package handlers

import (
	"net/http"

	"github.com/Dante983/go-blog/internal/auth"
)

// AddAuthData adds authentication status to template data
func AddAuthData(r *http.Request, data map[string]interface{}) {
	data["IsAuthenticated"] = auth.IsAuthenticated(r)
} 