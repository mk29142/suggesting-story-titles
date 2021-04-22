package benchmark_test

import (
	"github.com/mk29142/suggesting-story-titles/client"
	"github.com/mk29142/suggesting-story-titles/domain"
	"github.com/mk29142/suggesting-story-titles/workpool"
	"github.com/mk29142/suggesting-story-titles/workpool/internal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Benchmark", func() {
	var (
		fakeClient *internal.FakeClient
		tasks []workpool.Task

		coords                 domain.Metadata
		location               string
		higherConcurrencyValue int
    lowerConcurrencyValue  int
		numTasks               int
	)

	BeforeEach(func() {
		fakeClient = new(internal.FakeClient)
		coords = domain.Metadata{
			Latitude:  50.123,
			Longitude: 0.456,
		}
		location = "London"
		higherConcurrencyValue = 2
		lowerConcurrencyValue = 6
		numTasks = 1000000

		for i := 0; i < numTasks; i++  {
			tasks = append(tasks, workpool.NewTask(coords, fakeClient))
		}

		fakeClient.LocationReturns(client.Location{
			Latitude:  coords.Latitude,
			Longitude: coords.Longitude,
			Name:      location,
		}, nil)
	})

	Describe("Performance", func() {
		var (
			fasterPool workpool.Pool
			slowerPool workpool.Pool

			res1 []domain.Location
			errs1 []error

			res2 []domain.Location
			errs2 []error
		)

		BeforeEach(func() {
			fasterPool = workpool.New(tasks, higherConcurrencyValue)
			slowerPool = workpool.New(tasks, lowerConcurrencyValue)

			go func() {
				for out := range fasterPool.Output() {
					res1 = append(res1, domain.Location{
						Latitude:  out.Lat,
						Longitude: out.Long,
						Name:      out.Name,
					})
				}
			}()

			go func() {
				for err := range fasterPool.Errors() {
					errs2 = append(errs2, err)
				}
			}()

			go func() {
				for out := range slowerPool.Output() {
					res2 = append(res2, domain.Location{
						Latitude:  out.Lat,
						Longitude: out.Long,
						Name:      out.Name,
					})
				}
			}()

			go func() {
				for err := range slowerPool.Errors() {
					errs2 = append(errs1, err)
				}
			}()
		})

		Measure("time to process all tasks", func(b Benchmarker) {
			faster := b.Time("runtime", func() {
				fasterPool.Run()
				length := func() int {
					return len(res1)
				}
				Eventually(length, "3s", "1s").Should(Equal(numTasks))
			})

			slower := b.Time("runtime", func() {
				slowerPool.Run()
				length := func() int {
					return len(res1)
				}
				Eventually(length, "3s", "1s").Should(Equal(numTasks))
			})

			Expect(faster.Milliseconds()).To(BeNumerically("<", slower.Milliseconds()))
		}, 1)
	})
})
