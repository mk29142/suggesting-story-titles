package domain

import "time"

type Metadata struct {
  Timestamp time.Time
  Latitude  float64
  Longitude float64
}

type Location struct {
  Latitude  float64   `json:"lat"`
  Longitude float64   `json:"lng"`
  Timestamp time.Time `json:"timestamp"`
  Name      string    `json:"name"`
}

type Title struct {
  Latitude  float64   `json:"lat"`
  Longitude float64   `json:"lng"`
  Timestamp time.Time `json:"timestamp"`
  Name      string    `json:"name"`

  Title    string     `json:"title"`
}
