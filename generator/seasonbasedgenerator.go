package generator

import (
  "fmt"
  "github.com/mk29142/suggesting-story-titles/domain"
  "math/rand"
)

var seasonalSuggestions = map[string][]string {
  "winter": {"A chilly winter in %s", "%s in the winter"},
  "spring": {"A blooming spring in %s", "%s in the spring"},
  "summer": {"A beautiful summer in %s", "%s in the summer"},
  "autumn": {"A picturesque autumn in %s", "%s in the autumn"},
}

type SeasonBasedGenerator struct {}

func NewSeasonBasedGenerator() SeasonBasedGenerator {
  return SeasonBasedGenerator{}
}

func (s SeasonBasedGenerator) Generate(location domain.Location) Suggestion {
  month := location.Timestamp.Month()

  if month == 12 || month <= 2 {
    choices := seasonalSuggestions["winter"]
    title := fmt.Sprintf(choices[rand.Intn(len(choices))], location.Name)
    return Suggestion{
      Title:     title,
    }
  }

  if month >=3 && month <= 5 {
    choices := seasonalSuggestions["spring"]
    title := fmt.Sprintf(choices[rand.Intn(len(choices))], location.Name)
    return Suggestion{
      Title:     title,
    }
  }

  if month >=6 && month <= 8 {
    choices := seasonalSuggestions["summer"]
    title := fmt.Sprintf(choices[rand.Intn(len(choices))], location.Name)
    return Suggestion{
      Title:     title,
    }
  }

  if month >=9 && month <= 11 {
    choices := seasonalSuggestions["autumn"]
    title := fmt.Sprintf(choices[rand.Intn(len(choices))], location.Name)
    return Suggestion{
      Title:     title,
    }
  }

  return Suggestion{}
}