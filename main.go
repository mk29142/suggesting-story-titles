package main

import (
	"encoding/json"
	"fmt"
	"github.com/mk29142/suggesting-story-titles/client"
	"github.com/mk29142/suggesting-story-titles/csv"
	"github.com/mk29142/suggesting-story-titles/generator"
	"github.com/mk29142/suggesting-story-titles/orchestrator"
	"github.com/mk29142/suggesting-story-titles/service"
	"net/http"
	"os"
)

func main() {

	argLength := len(os.Args[1:])
	if argLength < 2 {
		fmt.Println("must provide csv file and valid api-token")
		os.Exit(2)
	}

	file := os.Args[1]
	apiToken := os.Args[2]

	reader := csv.NewReader()

	data, err := reader.Read(file)
	if err != nil {
		fmt.Println("error while reading csv: %w", err)
		os.Exit(2)
	}

	suggestor := service.NewSuggestor(generator.TimeBasedGenerator{}, generator.SeasonBasedGenerator{}, generator.GenericGenerator{})
	client := client.New(apiToken, &http.Client{})
	o := orchestrator.New(suggestor, client)

	titles := o.Titles(data)
	for _, title := range titles {
		out, _ := json.Marshal(title)
		fmt.Println(string(out))
	}
}
