package generator

import (
  "fmt"
  "github.com/mk29142/suggesting-story-titles/domain"
  "math/rand"
)

var genericSuggestions = []string{"A fun trip to %s", "Fun times in %s", "Take me back to %s", "Time well spent in %s"}

type GenericGenerator struct {}

func NewGenericGenerator() GenericGenerator {
  return GenericGenerator{}
}

func (g GenericGenerator) Generate(location domain.Location) Suggestion {
  title := fmt.Sprintf(genericSuggestions[rand.Intn(len(genericSuggestions))], location.Name)
  return Suggestion{
    Title:     title,
  }
}
