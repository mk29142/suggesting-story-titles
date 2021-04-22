package csv

import (
  "encoding/csv"
  "fmt"
  "github.com/mk29142/suggesting-story-titles/domain"
  "io"
  "os"
  "strconv"
  "time"
)

const format1 = "2006-01-02 15:04:05"
const format2 = "2006-01-02T15:04:05Z"

type Reader struct {
}

func NewReader() Reader {
  return Reader{}
}

func (r Reader) Read(file string) ([]domain.Metadata, error) {
  f, err := os.Open(file)
  if err != nil {
    return nil, err
  }
  defer f.Close()

  csvr := csv.NewReader(f)

  var data []domain.Metadata
  for {
    row, err := csvr.Read()
    if err != nil {
      if err == io.EOF {
        err = nil
      }
      return data, err
    }

    timeStamp, err := parse(row[0])
    if err != nil {
      return nil, fmt.Errorf("failed to parse csv: %w", err)
    }

    latitude, err := strconv.ParseFloat(row[1], 64)
    if err != nil {
      return nil, fmt.Errorf("failed to parse csv: %w", err)
    }

    longitude, err := strconv.ParseFloat(row[2], 64)
    if err != nil {
      return nil, fmt.Errorf("failed to parse csv: %w", err)
    }

    data = append(data, domain.Metadata{
      Timestamp: timeStamp,
      Latitude:  latitude,
      Longitude: longitude,
    })
  }
}

func parse(t string) (time.Time, error) {
  timeStamp, err := time.Parse(format1, t)
  if err == nil {
    return timeStamp, nil
  }

  timeStamp, err = time.Parse(format2, t)
  if err == nil {
    return timeStamp, nil
  }

  return time.Time{}, err
}
