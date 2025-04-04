package scraper

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/GregoryTomy/colorado-tax-rag/internal/models"
	"github.com/gocolly/colly/v2"
)

// parseTime tries different time formats to parse a timestamp
func parseTime(timeStr string) (time.Time, error) {
	formats := []string{
		time.RFC3339,
		"2006-01-02T15:04:05-07:00",
		"2006-01-02T15:04:05",
		"2006-01-02",
	}

	for _, format := range formats {
		t, err := time.Parse(format, timeStr)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse time: %s", timeStr)
}

// CollectSitemap parses a sitemap.xml from a URL and returns a list of URLs
func CollectSitemap(siteMapLink string) ([]models.SitemapURL, error) {
	collecter := colly.NewCollector(
		colly.UserAgent("ColoradoPropertyTaxScraper/1.0"),
	)

	var urls []models.SitemapURL

	// Extract URLs and metadata from sitemap
	collecter.OnXML("//urlset/url", func(xmlElement *colly.XMLElement) {
		var url models.SitemapURL

		url.Loc = xmlElement.ChildText("./loc")

		lastModStr := xmlElement.ChildText("./lastmod")
		if lastModStr != "" {
			parsedTime, error := parseTime(lastModStr)
			if error == nil {
				url.LastModified = parsedTime
			} else {
				log.Printf("Warning: could not parse lastmod time for %s: %v", url.Loc, error)
			}
		}
		// Extract change frequency if available
		url.ChangeFreq = xmlElement.ChildText("./changefreq")

		// Extract priority if available
		priorityStr := xmlElement.ChildText("./priority")
		if priorityStr != "" {
			priority, err := strconv.ParseFloat(priorityStr, 64)
			if err == nil {
				url.Priority = priority
			} else {
				log.Printf("Warning: could not parse priority for %s: %v", url.Loc, err)
			}
		}

		urls = append(urls, url)
	})

	collecter.OnError(func(response *colly.Response, err error) {
		log.Printf("Error fetching sitemap %s: %v", response.Request.URL, err)
	})

	err := collecter.Visit(siteMapLink)
	if err != nil {
		return nil, err
	}
	log.Printf("Extracted %d URLs from sitemap at %s", len(urls), siteMapLink)
	return urls, nil
}
