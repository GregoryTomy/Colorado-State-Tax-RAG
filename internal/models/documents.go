package models

import (
	"time"
)

// Document represents a scraped page from the Colorado Property Tax website
type Document struct {
	URL          string
	Tile         string
	Content      string
	Path         []string // Heirarchical path (volume/chapter/section)
	LastModified time.Time
	Metadata     map[string]string
	Timestamp    time.Time
}

// SitemapURL represents a URL from a sitemap with its metadata
type SitemapURL struct {
	Loc          string
	LastModified time.Time
	ChangeFreq   string
	Priority     float64
}
