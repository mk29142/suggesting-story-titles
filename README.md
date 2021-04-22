# suggesting-story-titles

This app takes in a csv file containing lat, long and timestamp 
and returns a suggested title.

#Input
```json
2020-03-30 14:12:19,40.728808,-73.996106
2020-03-30 14:20:10,40.728656,-73.998790
2020-03-30 14:32:02,40.727160,-73.996044
2020-03-30 14:41:18,40.725468,-73.995701
``` 

#Output
```json
{"lat":40.728808,"lng":-73.996106,"timestamp":"2020-03-30T14:12:19Z","name":"New York","title":"A blooming spring in New York"}
{"lat":40.728656,"lng":-73.99879,"timestamp":"2020-03-30T14:20:10Z","name":"New York","title":"Time well spent in New York"}
{"lat":40.72716,"lng":-73.996044,"timestamp":"2020-03-30T14:32:02Z","name":"New York","title":"Time well spent in New York"}
{"lat":40.725468,"lng":-73.995701,"timestamp":"2020-03-30T14:41:18Z","name":"New York","title":"A fun afternoon in New York"}
```

#How to use
```bash
make build
./suggesting-story-titles <path to csv> <api-token>
```

#Setup
```bash
make init
```

Unit Tests
```bash
make test-unit
```

Benchmark Tests
```bash
make test-benchmark
```

#Design Decisions

###Problem: What is the best way to get the locations from all the coordinates?

###Solution 1: 
Map box provides batching, so you can batch all the `lat,long` combos as query params.

Pros: 
* Reduces network traffic, only need one request.

Cons: 
* Single point of failure, if request fails, all photos won't get their results.
* Time to execute api-results will get longer and longer given size of csv.

###Solution 2

Use a worker pool so that each coordinate is a task and tasks are processed concurrently. 

Pros:
 * Scales well with large albums, so faster response to users
 * If one image in the album fails, it doesn't cause the entire request to fail.

Cons: 
 * Code is slightly more complex
 * Higher network traffic

I choose to do solution 2 as I felt that it provided a better user experience in terms
of speed and better scalability.

###Abstractions
* I decoupled the logic of getting the location of the lat long combos from the worker 
   pool, so that it leaves room for extension, as it will make it easier in the future if 
   we want to process the inputs differently. 

* For suggestions I use multiple different generators, time based, season based and generic.
  The `suggestor` takes in an interface so won't require code changes if you want to create more
  generators. As you would just need to follow the interface contract of a generator and just pass
  then into the constructor. 
     
#TODO
 * In a production setting I would pass in the api-token using a config file. 
  
 * I haven't added a system test at the moment. For a system test, it would mean using something like 
   [direnv](https://direnv.net/) to set a profile file with a valid api-token so that the system tests will
   test the full end-to-end flow that will actually call the api. 
 
 * The benchmark test is a bit flakey due to scheduling. Ideally I would take the average 
 of multiple runs rather than just a single run to reduce flakey results when it comes to 
 scheduling go routines.
 
 * In a production setting I would create a logger and pass in the context. This would
   allow logging with structured arguments that would store the input lat and longs. 
   
 * This code only handles errors by printing then to stdErr, ideally we would want to handle errors
   in a better way.

 * I think the anti-corruption layers can be improved, i.e better decoupling on models between the layers
   so that each layer has its own set of models.
 
#Notes
* Uses the [Mapbox](https://www.mapbox.com/) api to perform reverse geo lookup. 

* This project uses [ginkgo](https://github.com/onsi/ginkgo) for tests.
* This project used [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter) to generate mocks for 
  any interfaces annotated with the following:
```golang
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o internal/fake_client.go . Client
type Client interface {
  Location(coordinates domain.Metadata) (client.Location, error)
}
```
The mocks are stored in the `internal` directory of a package. 

* I don't have much experience with concurrency in go in a production environment 
so the way I have implemented the worker pool pattern may not be the most idiomatic go.
