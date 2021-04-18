package workpool_test

import (
	"fmt"
	"github.com/mk29142/suggesting-story-titles/client"
	"github.com/mk29142/suggesting-story-titles/domain"
	"github.com/mk29142/suggesting-story-titles/workpool"
	"github.com/mk29142/suggesting-story-titles/workpool/internal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pool", func() {
	var (
		fakeClient *internal.FakeClient
		tasks []workpool.Task

		coords domain.Coordinates
		postcode string
		concurrency int
	)

	BeforeEach(func() {
		fakeClient = new(internal.FakeClient)
		coords = domain.Coordinates{
			Latitude:  50.123,
			Longitude: 0.456,
		}
		postcode = "SS16 5HE"
		concurrency = 2

		task1 := workpool.NewTask(coords, fakeClient)
		task2 := workpool.NewTask(coords, fakeClient)
		task3 := workpool.NewTask(coords, fakeClient)

		tasks = []workpool.Task{task1, task2, task3}

		fakeClient.PostcodeReturns(client.LatLongPostcode{
			Latitude:  coords.Latitude,
			Longitude: coords.Longitude,
			Postcode:  postcode,
		}, nil)
	})

	Describe("Success", func() {
		var (
			pool workpool.Pool
		)

		BeforeEach(func() {
			pool = workpool.New(tasks, concurrency)
		})

		When("success", func() {
			var (
				res []domain.Postcode
				errs []error
			)

			BeforeEach(func() {
				go func() {
					for out := range pool.Output() {
						res = append(res, domain.Postcode{
							Latitude:  out.Lat,
							Longitude: out.Long,
							Postcode:  out.PostCode,
						})
					}
				}()

				go func() {
					for err := range pool.Errors() {
						errs = append(errs, err)
					}
				}()
			})

			It("adds to output channel", func() {
				pool.Run()

				expect := []domain.Postcode {
					{
						Latitude:  coords.Latitude,
						Longitude: coords.Longitude,
						Postcode:  postcode,
					},
					{
						Latitude:  coords.Latitude,
						Longitude: coords.Longitude,
						Postcode:  postcode,
					},
					{
						Latitude:  coords.Latitude,
						Longitude: coords.Longitude,
						Postcode:  postcode,
					},
				}

				length := func() int {
					return len(res)
				}

				Expect(errs).To(BeEmpty())
				Eventually(length, "3s", "1s").Should(Equal(3))
				Expect(res).To(Equal(expect))
			})
		})

		When("Error", func() {
			var (
				res []domain.Postcode
				errs []error
			)

			BeforeEach(func() {
				fakeClient.PostcodeReturns(client.LatLongPostcode{}, fmt.Errorf("something went wrong"))

				go func() {
					for out := range pool.Output() {
						res = append(res, domain.Postcode{
							Latitude:  out.Lat,
							Longitude: out.Long,
							Postcode:  out.PostCode,
						})
					}
				}()

				go func() {
					for err := range pool.Errors() {
						errs = append(errs, err)
					}
				}()
			})

			It("adds to error channel", func() {
				pool.Run()

				length := func() int {
					return len(errs)
				}

				Expect(res).To(BeEmpty())
				Eventually(length, "3s", "1s").Should(Equal(3))
			})
		})
	})
})
