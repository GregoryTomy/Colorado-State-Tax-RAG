package vectordb

import (
	"context"

	"github.com/GregoryTomy/colorado-tax-rag/internal/models"
	"github.com/philippgille/chromem-go"
)

type EmbeddingGenerator func(ctx context.Context, text string) ([]float32, error)

// ChromemStore provides vector storage and retrieval using chromem-go
type ChromemStore struct {
	db             *chromem.DB
	collection     *chromem.Collection
	embedGenerator EmbeddingGenerator
}

// NewChromemStore creates a new ChromemStore
func NewChromemStore(collectionName string, persistencePath string, embedder EmbeddingGenerator) (*ChromemStore, error) {
}

// AddDocuments method adds documents to the vector store
func (store *ChromemStore) AddDocuments(ctx *context.Context, documents []models.Document) error {
	panic("")
}

// Query method searches and
