package domain

import "time"

// Post represents a blog post
type Post struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Slug        string    `json:"slug"`
	Author      string    `json:"author"`
	IsPublished bool      `json:"is_published"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreatePostInput is used when creating a new post
type CreatePostInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Slug    string `json:"slug"`
	Author  string `json:"author"`
}

// UpdatePostInput is used when updating a post
type UpdatePostInput struct {
	Title       *string `json:"title,omitempty"`
	Content     *string `json:"content,omitempty"`
	Slug        *string `json:"slug,omitempty"`
	IsPublished *bool   `json:"is_published,omitempty"`
}
