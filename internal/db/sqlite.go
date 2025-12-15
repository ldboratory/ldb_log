Errorf("failed to run migrations: %w", err)
	}

	return db, nil
}

// runMigrations creates tables if they don't exist
func runMigrations(db *sql.DB) error {
	migrations := []string{
		// Posts table
		`CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			content TEXT NOT NULL,
			slug TEXT UNIQUE NOT NULL,
			is_published BOOLEAN DEFAULT 0,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,

		// Comments table (anonymous)
		`CREATE TABLE IF NOT EXISTS comments (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			post_id INTEGER NOT NULL,
			author_name TEXT NOT NULL,
			author_email TEXT,
			content TEXT NOT NULL,
			parent_id INTEGER,
			created_at DATETIME DEpackage db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

// InitSQLite initializes SQLite database connection
func InitSQLite(dsn string) (*sql.DB, error) {
	// Create data directory if it doesn't exist
	dir := filepath.Dir(dsn)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create data directory: %w", err)
	}

	// Open database connection
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Run migrations
	if err := runMigrations(db); err != nil {
		return nil, fmt.FAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
			FOREIGN KEY (parent_id) REFERENCES comments(id) ON DELETE CASCADE
		)`,

		// Create indexes for better query performance
		`CREATE INDEX IF NOT EXISTS idx_posts_slug ON posts(slug)`,
		`CREATE INDEX IF NOT EXISTS idx_posts_created_at ON posts(created_at DESC)`,
		`CREATE INDEX IF NOT EXISTS idx_comments_post_id ON comments(post_id)`,
	}

	for _, migration := range migrations {
		if _, err := db.Exec(migration); err != nil {
			return fmt.Errorf("migration failed: %w", err)
		}
	}

	return nil
}
