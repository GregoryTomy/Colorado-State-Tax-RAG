package vectordb

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/GregoryTomy/colorado-tax-rag/internal/models"
)

// SQLLoader loads documents from SQL into the vector store.
type SQLLoader struct {
	db       *sql.DB
	vectorDB *ChromemStore
}

// NewSQLLoader creates a new loader for loading documents from SQLite to vector store.
func NewSQLLoader(db *sql.DB, vectorDB *ChromemStore) *SQLLoader {
	return &SQLLoader{
		db,
		vectorDB,
	}
}

// LoadAllDocuments loads all documents from SQL to vector store.
func (loader *SQLLoader) LoadAllDocuments(ctx context.Context) error {
	log.Println("Loading documents from SQL database...")

	rows, err := loader.db.QueryContext(ctx, `
    SELECT url, title, content
    FROM documents
    `)
	if err != nil {
		return fmt.Errorf("failed to query documents from SQL database: %w", err)
	}

	defer rows.Close()

	// TODO: Improve integration to directly work with Chromem Documents instead of our struct
	var documents []models.Document
	for rows.Next() {
		var doc models.Document

		err := rows.Scan(
			&doc.URL,
			&doc.Title,
			&doc.Content,
		)
		if err != nil {
			return fmt.Errorf("failed to scan document row: %w", err)
		}
		documents = append(documents, doc)

		if err = rows.Err(); err != nil {
			return fmt.Errorf("error during rows iteration: %w", err)
		}
	}
	log.Printf("Found %d documents in SQL database.", len(documents))

	if len(documents) > 0 {
		return loader.vectorDB.AddDocuments(ctx, documents)
	}
	return nil
}
