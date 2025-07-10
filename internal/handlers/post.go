package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/Dante983/go-blog/internal/models"
	"github.com/Dante983/go-blog/internal/views"
)

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	// Parse page number from query params
	page := 1
	if p := r.URL.Query().Get("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			page = parsed
		}
	}

	// Set pagination params
	limit := 10
	offset := (page - 1) * limit

	// Get posts
	posts, err := models.GetAllPosts(limit+1, offset) // Get one extra to check if there are more
	if err != nil {
		http.Error(w, "Failed to retrieve posts", http.StatusInternalServerError)
		return
	}

	// Check if there are more posts
	hasMore := len(posts) > limit
	if hasMore {
		posts = posts[:limit] // Remove the extra post
	}

	// Prepare template data
	data := map[string]interface{}{
		"Title":    "All Posts - Nikola's Blog",
		"Posts":    posts,
		"HasMore":  hasMore,
		"NextPage": page + 1,
	}
	
	AddAuthData(r, data)

	// Check if this is an HTMX request (for load more functionality)
	if r.Header.Get("HX-Request") == "true" && page > 1 {
		// For HTMX requests, only render the posts partial
		for _, post := range posts {
			fmt.Fprintf(w, `<article class="post-preview">
				<h3><a href="/posts/%s">%s</a></h3>
				<time>%s</time>
				<p>%s</p>
				<a href="/posts/%s">Read more →</a>
			</article>
			<hr>`, post.Slug, post.Title, post.CreatedAt.Format("January 2, 2006"), post.Summary(), post.Slug)
		}
		return
	}

	// Render full page
	err = views.Templates.ExecuteTemplate(w, "posts", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SinglePostHandler(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	
	// Get post by slug
	post, err := models.GetPostBySlug(slug)
	if err != nil {
		http.Error(w, "Failed to retrieve post", http.StatusInternalServerError)
		return
	}
	
	if post == nil {
		http.NotFound(w, r)
		return
	}
	
	// Prepare data with title
	data := map[string]interface{}{
		"Title":      post.Title + " - Nikola's Blog",
		"ID":         post.ID,
		"PostTitle":  post.Title,
		"Slug":       post.Slug,
		"Content":    post.Content,
		"ContentHTML": post.ContentHTML(),
		"CreatedAt":  post.CreatedAt,
		"UpdatedAt":  post.UpdatedAt,
	}
	
	AddAuthData(r, data)
	
	// Render post
	err = views.Templates.ExecuteTemplate(w, "post", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SnippetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<div style="border: 1px solid #00ff00; padding: 10px; background: rgba(0,255,0,0.1);">
		<p style="color: #00ff00; margin: 0;">> CONNECTION ESTABLISHED</p>
		<p style="color: #00ff00; margin: 0;">> HTMX PROTOCOL: <span style="color: #00ffff;">ACTIVE</span></p>
		<p style="color: #00ff00; margin: 0;">> LATENCY: <span style="color: #00ffff;">12ms</span></p>
	</div>`))
}
