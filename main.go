package main

import (
	"log"

	"github.com/GregoryTomy/colorado-tax-rag/internal/database"
	"github.com/GregoryTomy/colorado-tax-rag/internal/scraper"
)

var testURLs = []string{
	"https://arl.colorado.gov/chapter-1-statutory-and-case-law-references",
	// "https://arl.colorado.gov/chapter-2-appraisal-process-economic-areas-and-the-approaches-to",
	// "https://arl.colorado.gov/chapter-3-sales-confirmation-and-stratification",
}

func main() {
	store, err := database.NewSQLiteStore("./document-store")
	if err != nil {
		log.Fatalf("Failed to create new SQLite Store: %v", err)
	}

	defer store.Close()

	for _, url := range testURLs {
		document, err := scraper.ScrapeURL(url)
		if err != nil {
			log.Printf("Error scraping %s: %v", url, err)
			continue
		}

		if err := store.StoreDocument(document); err != nil {
			log.Printf("Error stroing document: %s", document.URL)
		}
	}
}

// context := context.Background()

// log.Println("Setting up chromem-go...")

// db, err := chromem.NewPersistentDB("./tax-raw", false)
// if err != nil {
// 	panic(err)
// }
//
// // create the collection for documents
// collection, err := db.GetOrCreateCollection("tax-1", nil, nil)
// if err != nil {
// 	log.Fatalf("Failed to create collect %v", err)
// }
//
// create a variable named documents that is an empty slice
// that will hold pointers to Document objects
// var chromemDocuments []chromem.Document
//
// for _, url := range testURLs {
// 	document, err := scraper.ScrapeURL(url)
// 	if err != nil {
// 		log.Printf("Error scraping %s: %v", url, err)
// 		continue
// 	}
//
// 	chromemDocument := chromem.Document{
// 		ID:       document.URL,
// 		Metadata: make(map[string]string),
// 		Content:  document.Content,
// 	}
//
// 	chromemDocuments = append(chromemDocuments, chromemDocument)
// }
//
// // if collection.Count() == 0 {
// //
// // }
//
// if len(chromemDocuments) > 0 {
// 	log.Printf("Adding %d documents to collection...", len(chromemDocuments))
// }
//
// err = collection.AddDocuments(context, chromemDocuments, 2)
// if err != nil {
// 	log.Fatalf("Failed to add documents to collection: %v", err)
// }
//
// log.Printf("Successfully added documents to collection")
