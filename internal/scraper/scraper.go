package scraper

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/GregoryTomy/colorado-tax-rag/internal/models"
	"github.com/gocolly/colly/v2"
)

func SrapeURL(siteURL string) (*models.Document, error) {
	collector := colly.NewCollector(
		colly.UserAgent("ColoradoPropertyTaxScraper/1.0"),
		colly.MaxDepth(1),
	)

	// create an instance of Document to store scraped data
	// and assigns its memory address to the document variable
	document := &models.Document{
		URL:       siteURL,
		Timestamp: time.Now(),
		Metadata:  make(map[string]string), // initalize an empty Metadata map
	}

	collector.OnHTML("div.region.region-header h1", func(e *colly.HTMLElement) {
		document.Title = strings.TrimSpace(e.Text)
		log.Printf("Found title: %s", document.Title)
	})

	collector.OnHTML("main.main-container section.col-sm-9", func(e *colly.HTMLElement) {
		document.Content = e.Text
		log.Printf("Scraped page content raw.")
	})
	// Handle errors
	collector.OnError(func(r *colly.Response, err error) {
		log.Printf("Error scraping %s: %v", r.Request.URL, err)
	})

	// Visit the URL - this is where the HTTP request is made
	err := collector.Visit(siteURL)
	if err != nil {
		return nil, fmt.Errorf("failed to visit URL %s: %w", siteURL, err)
	}
	return document, nil
}
