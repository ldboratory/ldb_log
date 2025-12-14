package domain

import "time"

// UserRole represents user roles
type UserRole string

const (
	RoleAdmin UserRole = "admin"
	RoleUser  UserRole = "user"
)

// User represents a user (mainly for admin)
type User struct {
	ID           int64     `json:"id"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"-"` // never expose in JSON
	Role         UserRole  `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
}

// LoginInput is used for authentication
type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
