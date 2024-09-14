package routes

import (
	"net/http"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")

		_, err := db.Exec("DELETE FROM posts WHERE id = ?", id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/blog", http.StatusSeeOther)
		return
	}
}
