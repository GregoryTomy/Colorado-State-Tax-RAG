package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/GregoryTomy/colorado-tax-rag/internal/vectordb"
	_ "github.com/mattn/go-sqlite3"
	"github.com/philippgille/chromem-go"
)

const (
	embeddingModel = "nomic-embed-text"
	ollamaIP       = "http://127.0.0.1:11434/api"
)

func main() {
	ctx := context.Background()

	log.Println("Starting SQL to Vector DB transfer test...")
	log.Printf("Embedding Model: %s", embeddingModel)

	sqlDB, err := sql.Open("sqlite3", "./document-store")
	if err != nil {
		log.Fatalf("Failed to open SQLite database: %v", err)
	}

	defer sqlDB.Close()

	embeddingFunction := chromem.NewEmbeddingFuncOllama(
		embeddingModel,
		ollamaIP,
	)

	vectorStore, err := vectordb.NewChromemStore(
		"colorado-tax",
		"./vector-store",
		embeddingFunction,
	)
	if err != nil {
		log.Fatalf("Failed to create vector store: %v", err)
	}

	log.Println("Successfully created vector store.")

	loader := vectordb.NewSQLLoader(
		sqlDB,
		vectorStore,
	)

	log.Println("Starting document transfer with embedding...")
	if err := loader.LoadAllDocuments(ctx); err != nil {
		log.Fatalf("Failed to load documents to vector store: %v", err)
	}

	log.Println("Successfully loaded documents from SQL to vector store.")
}
