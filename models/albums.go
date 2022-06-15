package models

type Albums struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	Singer      float64 `json:"singer"`
	ReleaseYear string  `json:"release_year"`
}
