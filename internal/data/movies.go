package data

import "time"

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // Use the "-" tag to omit this field from JSON output
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"` // Omit if zero value
	Runtime   Runtime   `json:"runtime,omitempty,string"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}
