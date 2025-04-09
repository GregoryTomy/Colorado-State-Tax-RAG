package vectordb

import (
	"context"
	"fmt"
	"log"
	"runtime"

	"github.com/GregoryTomy/colorado-tax-rag/internal/models"
	"github.com/philippgille/chromem-go"
)

// ChromemStore provides vector storage and retrieval using chromem-go
type ChromemStore struct {
	db         *chromem.DB
	collection *chromem.Collection
}

// NewChromemStore creates a new ChromemStore
func NewChromemStore(collectionName string, persistencePath string, embedFunc chromem.EmbeddingFunc) (*ChromemStore, error) {
	if persistencePath == "" {
		persistencePath = "./chromem-db"
	}

	log.Println("Setting up chromem-go")

	// TODO: check behavior if DB already exists
	db, err := chromem.NewPersistentDB(persistencePath, false)
	if err != nil {
		return nil, fmt.Errorf("failed to create persistent chromem DB: %w", err)
	}

	collection, err := db.GetOrCreateCollection(collectionName, nil, embedFunc)
	if err != nil {
		return nil, fmt.Errorf("failed to get or create collection: %w", err)
	}

	return &ChromemStore{
		db,
		collection,
	}, nil
}

// AddDocuments method adds documents to the vector store
func (store *ChromemStore) AddDocuments(ctx context.Context, documents []models.Document) error {
	if len(documents) == 0 {
		return nil
	}

	// convert our model.Documents to chromem.Document
	chromemDocs := make([]chromem.Document, 0, len(documents))

	// TODO: Adding prefix as search document to improve retriveal
	// TODO: Adding metadata
	for _, document := range documents {
		chromemDoc := chromem.Document{
			ID:      document.URL,
			Content: document.Content,
		}

		chromemDocs = append(chromemDocs, chromemDoc)
	}

	numCPUs := runtime.NumCPU()

	err := store.collection.AddDocuments(ctx, chromemDocs, numCPUs)
	if err != nil {
		return fmt.Errorf("failed to add documents to vector store: %w", err)
	}

	log.Printf("Added %d documents to vector store", len(documents))
	return nil
}

// TODO: Conditional filtering with metadata.
// QuerySimilar method finds documents similar to the provided query text
func (store *ChromemStore) QuerySimilar(ctx context.Context, queryText string, numResults int) ([]chromem.Result, error) {
	results, err := store.collection.Query(ctx, queryText, numResults, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to query database: %w", err)
	}

	return results, nil
}

// CountDocuments method counts the number of documents in the collection
func (store *ChromemStore) Count() int {
	return store.collection.Count()
}

// Reset method removes all documents from collection.
func (store *ChromemStore) Reset() error {
	return store.db.Reset()
}
