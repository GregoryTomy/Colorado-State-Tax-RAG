package main

import (
	"fmt"
	"log"

	"github.com/GregoryTomy/colorado-tax-rag/internal/scraper"
)

func main() {
	sitemapLink := "https://arl.colorado.gov/sitemap.xml"

	urls, err := scraper.CollectSitemap(sitemapLink)
	if err != nil {
		log.Fatalf("Error collecting sitemap: %v", err)
	}

	fmt.Printf("Total urls extracted %d\n", len(urls))
	for _, url := range urls {
		fmt.Printf("URL: %s\n", url.Loc)
		fmt.Printf("Last Modified: %v\n", url.LastModified)
		fmt.Printf("Change Frequency: %s\n", url.ChangeFreq)
		fmt.Printf("Priority: %f\n\n", url.Priority)
	}
}
