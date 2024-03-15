package data

import (
	"encoding/json"
	"fmt"
	"time"
)

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // Use the - directive
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`   // Add the omitempty directive
	Runtime   Runtime   `json:"-"`                // Add the omitempty directive
	Genres    []string  `json:"genres,omitempty"` // Add the omitempty directive
	Version   int32     `json:"version"`
}

func (m Movie) MarshalJSON() ([]byte, error) {
	// Declare a variable to hold the custom runtime string (this will be the empty // string "" by default).
	var runtime string
	if m.Runtime != 0 {
		runtime = fmt.Sprintf("%d mins", m.Runtime)
	}
	type MovieAlias Movie
	aux := struct {
		MovieAlias
		Runtime string `json:"runtime,omitempty"`
	}{
		MovieAlias: MovieAlias(m),
		Runtime:    runtime,
	}
	return json.Marshal(aux)
}
