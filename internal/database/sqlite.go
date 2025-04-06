package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/GregoryTomy/colorado-tax-rag/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

// SQLiteStore manages document storage using SQLite
type SQLiteStore struct {
	db *sql.DB
}

// NewSQLiteStore creates and initializes a new SQLite document store
func NewSQLiteStore(dbPath string) (*SQLiteStore, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open SQLite database: %w", err)
	}

	store := &SQLiteStore{db}
	if err := store.initSchema(); err != nil {
		db.Close()
		return nil, err
	}

	return store, nil
}

// creating methods for SQLite Store
// //initSchema creates the database schema if it doesn't exist.
func (store *SQLiteStore) initSchema() error {
	_, err := store.db.Exec(`
    CREATE TABLE IF NOT EXISTS documents (
    id TEXT PRIMARY KEY,
    url TEXT NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    last_modified TIMESTAMP,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
    `)
	if err != nil {
		return fmt.Errorf("failed to create schema: %w", err)
	}

	return nil
}

func (store *SQLiteStore) StoreDocument(document *models.Document) error {
	transaction, err := store.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// pattern to safely rollback databse transaction if anything goes wrong
	defer func() {
		if err != nil {
			transaction.Rollback()
		}
	}()

	docID := document.URL

	_, err = transaction.Exec(
		"INSERT OR REPLACE INTO documents (id, url, title, content, last_modified, timestamp) VALUES (?, ?, ?, ?, ?, ?)",
		docID,
		document.URL,
		document.Title,
		document.Content,
		document.LastModified,
		time.Now(),
	)
	if err != nil {
		return fmt.Errorf("failed to insert document: %w", err)
	}

	return transaction.Commit()
}

// Close closes the database connection
func (store *SQLiteStore) Close() error {
	return store.db.Close()
}
