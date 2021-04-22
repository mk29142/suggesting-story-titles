package generator

import (
  "fmt"
  "github.com/mk29142/suggesting-story-titles/domain"
  "math/rand"
)

var timeSuggestions = map[string][]string {
  "morning": {"A lovely morning in %s", "A peaceful morning in %s", "%s in the morning"},
  "afternoon": {"Enjoying the afternoon in %s", "A fun afternoon in %s", "%s in the afternoon"},
  "night": {"Evening in %s", "A quiet evening walk in %s", "A blissful night spent in %s"},
}

type TimeBasedGenerator struct {

}

func NewTimeBasedGenerator() TimeBasedGenerator {
  return TimeBasedGenerator{}
}

func (t TimeBasedGenerator) Generate(location domain.Location) Suggestion {
   ts := location.Timestamp.UTC()

   if ts.Hour() >= 0 && ts.Hour() < 12 {
      choices := timeSuggestions["morning"]
      title := fmt.Sprintf(choices[rand.Intn(len(choices))], location.Name)
      return Suggestion{
        Title:     title,
      }
   }

  if ts.Hour() >= 12 && ts.Hour() < 18 {
    choices := timeSuggestions["afternoon"]
    title := fmt.Sprintf(choices[rand.Intn(len(choices))], location.Name)
    return Suggestion{
      Title:     title,
    }
  }

  if ts.Hour() >= 18 && ts.Hour() < 24 {
    choices := timeSuggestions["night"]
    title := fmt.Sprintf(choices[rand.Intn(len(choices))], location.Name)
    return Suggestion{
      Title:     title,
    }
  }

  return Suggestion{}
}
