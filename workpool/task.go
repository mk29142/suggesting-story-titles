package workpool

import (
  "fmt"
  "github.com/mk29142/suggesting-story-titles/client"
  "github.com/mk29142/suggesting-story-titles/domain"
  "time"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o internal/fake_client.go . Client
type Client interface {
  Location(coordinates domain.Metadata) (client.Location, error)
}

type Task struct {
  Coordinates domain.Metadata
  GeoCoder    Client
}

type Location struct {
  Lat float64
  Long float64
  Name string
  Timestamp time.Time
}

func NewTask(latLong domain.Metadata, geocoder Client) Task {
  return Task{
    Coordinates: latLong,
    GeoCoder:    geocoder,
  }
}

func (task Task) Process() (Location, error) {
  res, err := task.GeoCoder.Location(task.Coordinates)
  if err != nil {
    return Location{},
    domain.NewTaskError(task.Coordinates.Latitude,
      task.Coordinates.Longitude,
      fmt.Errorf("failure process task: %w", err))
    }

  return Location{
    Lat:  res.Latitude,
    Long: res.Longitude,
    Name: res.Name,
    Timestamp: res.Timestamp,
  }, nil
}