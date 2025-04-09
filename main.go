package main

import (
	"log"

	"github.com/GregoryTomy/colorado-tax-rag/internal/database"
	"github.com/GregoryTomy/colorado-tax-rag/internal/scraper"
)

var testURLs = []string{
	"https://arl.colorado.gov/chapter-1-statutory-and-case-law-references",
	"https://arl.colorado.gov/chapter-2-appraisal-process-economic-areas-and-the-approaches-to",
	"https://arl.colorado.gov/chapter-3-sales-confirmation-and-stratification",
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
