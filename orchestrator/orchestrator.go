package orchestrator

import (
  "fmt"
  "github.com/mk29142/suggesting-story-titles/client"
  "github.com/mk29142/suggesting-story-titles/domain"
  "github.com/mk29142/suggesting-story-titles/service"
  "github.com/mk29142/suggesting-story-titles/workpool"
  "os"
  "sync"
)

const POOLSIZE = 5

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o internal/fake_client.go . Client
type Client interface {
  Location(coordinates domain.Metadata) (client.Location, error)
}

//counterfeiter:generate -o internal/fake_suggestor.go . Suggestor
type Suggestor interface {
  Suggestions(locations []domain.Location) []service.Suggestion
}

type Orchestrator struct {
  Suggestor Suggestor
  GeoCoder  Client
}

func New(service Suggestor, geoCoder Client) Orchestrator {
  return Orchestrator{
    Suggestor: service,
    GeoCoder:  geoCoder,
  }
}

func (o Orchestrator) Titles(data []domain.Metadata) []domain.Title {
  locations := o.locations(data)

  ss := o.Suggestor.Suggestions(locations)

  var output []domain.Title
  for _, s := range ss {
    output = append(output, domain.Title{
      Latitude:  s.Latitude,
      Longitude: s.Longitude,
      Timestamp: s.Timestamp,
      Name:      s.Name,
      Title:     s.Title,
    })
  }

  return output
}

func (o Orchestrator) locations(data []domain.Metadata) []domain.Location {
  var tasks []workpool.Task
  for _, d := range data {
    tasks = append(tasks, workpool.NewTask(d, o.GeoCoder))
  }

  pool := workpool.New(tasks, POOLSIZE)

  var wg sync.WaitGroup
  wg.Add(1)

  var locations []domain.Location
  go func() {
    defer wg.Done()
    for res := range pool.Output() {
      l := domain.Location{
        Latitude:  res.Lat,
        Longitude: res.Long,
        Name:      res.Name,
        Timestamp: res.Timestamp,
      }

      locations = append(locations, l)
    }
  }()

  go func() {
    for err := range pool.Errors() {
      fmt.Fprintln(os.Stderr, err)
    }
  }()

  pool.Run()
  pool.Stop()

  wg.Wait()

  return locations
}




