package models

import (
	"database/sql"
	"html/template"
	"time"

	"github.com/Dante983/go-blog/internal/db"
	"github.com/Dante983/go-blog/internal/utils"
)

type Post struct {
	ID         int
	Title      string
	Slug       string
	Content    string
	CategoryID *int
	Category   *Category
	Tags       []Tag
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// ContentHTML returns the content as HTML from Markdown
func (p *Post) ContentHTML() template.HTML {
	return utils.MarkdownToHTML(p.Content)
}

// Summary returns a summary of the post content
func (p *Post) Summary() string {
	return utils.GetSummary(p.Content, 200)
}

// GetAllPosts retrieves all posts with pagination
func GetAllPosts(limit, offset int) ([]Post, error) {
	query := `
		SELECT id, title, slug, content, category_id, created_at, updated_at 
		FROM posts 
		ORDER BY created_at DESC 
		LIMIT $1 OFFSET $2
	`

	rows, err := db.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Slug, &post.Content,
			&post.CategoryID, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// GetPostBySlug retrieves a single post by its slug
func GetPostBySlug(slug string) (*Post, error) {
	query := `
		SELECT id, title, slug, content, category_id, created_at, updated_at 
		FROM posts 
		WHERE slug = $1
	`

	var post Post
	err := db.DB.QueryRow(query, slug).Scan(
		&post.ID, &post.Title, &post.Slug, &post.Content,
		&post.CategoryID, &post.CreatedAt, &post.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &post, nil
}

// CountPosts returns the total number of posts
func CountPosts() (int, error) {
	var count int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM posts").Scan(&count)
	return count, err
}
