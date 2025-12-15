package domain

import "time"

// Comment represents a blog post comment
type Comment struct {
	ID          int64     `json:"id"`
	PostID      int64     `json:"post_id"`
	AuthorName  string    `json:"author_name"`
	AuthorEmail string    `json:"author_email,omitempty"` // optional
	Content     string    `json:"content"`
	ParentID    *int64    `json:"parent_id,omitempty"` // for nested comments (future)
	CreatedAt   time.Time `json:"created_at"`
}

// CreateCommentInput is used when creating a new comment
type CreateCommentInput struct {
	PostID      int64  `json:"post_id"`
	AuthorName  string `json:"author_name"`
	AuthorEmail string `json:"author_email,omitempty"` // optional
	Content     string `json:"content"`
	ParentID    *int64 `json:"parent_id,omitempty"`
}
