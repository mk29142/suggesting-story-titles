package service

import (
  "github.com/mk29142/suggesting-story-titles/domain"
  "github.com/mk29142/suggesting-story-titles/generator"
  "math/rand"
  "time"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o internal/fake_generator.go . Generator
type Generator interface {
  Generate(location domain.Location) generator.Suggestion
}

type Service struct {
   Generators []Generator
}

type Suggestion struct {
  Latitude  float64
  Longitude float64
  Timestamp time.Time
  Name      string

  Title    string
}

func NewSuggestor(generators ...Generator) Service {
  return Service{Generators:generators}
}

func (s Service) Suggestions(locations []domain.Location) []Suggestion {
  var suggestions []Suggestion

  for _, l := range locations {
    randIndex := rand.Intn(len(s.Generators))
    randomGenerator := s.Generators[randIndex]
    title := randomGenerator.Generate(l)

    suggestions = append(suggestions, Suggestion{
      Latitude:  l.Latitude,
      Longitude: l.Longitude,
      Timestamp: l.Timestamp,
      Name:      l.Name,
      Title:     title.Title,
    })
  }

  return suggestions
}
