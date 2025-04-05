package main

import (
	"fmt"

	"github.com/GregoryTomy/colorado-tax-rag/internal/scraper"
)

func main() {
	url := "https://arl.colorado.gov/chapter-8-statistical-measurements"
	document, err := scraper.SrapeURL(url)
	if err != nil {
		fmt.Println("There was an error")
	}
	fmt.Println(document)
}
