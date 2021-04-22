package workpool_test

import (
  "fmt"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  "github.com/mk29142/suggesting-story-titles/client"
  "github.com/mk29142/suggesting-story-titles/domain"
  "github.com/mk29142/suggesting-story-titles/workpool"
  "github.com/mk29142/suggesting-story-titles/workpool/internal"
)

var _ = Describe("Task", func() {
	var (
		clientFake *internal.FakeClient
		t          workpool.Task

		coords = domain.Metadata{
			Latitude:  50.123,
			Longitude: 0.456,
		}

		location = "London"
	)

	BeforeEach(func() {
		clientFake = new(internal.FakeClient)
		t = workpool.NewTask(coords, clientFake)
	})

	Describe("Process", func() {
		var (
			result workpool.Location
			err    error
		)

		JustBeforeEach(func() {
			result, err = t.Process()
		})

		When("success", func() {
			BeforeEach(func() {
				resp := client.Location{
					Latitude:  coords.Latitude,
					Longitude: coords.Longitude,
					Name:      location,
				}

				clientFake.LocationReturns(resp, nil)
			})

			It("does not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("calls client", func() {
				Expect(clientFake.LocationCallCount()).To(Equal(1))
				got := clientFake.LocationArgsForCall(0)
				Expect(got).To(Equal(coords))
			})

			It("returns correct result", func() {
				Expect(result).To(Equal(workpool.Location{
					Lat:  coords.Latitude,
					Long: coords.Longitude,
					Name: location,
				}))
			})
		})

		When("client fails to get location", func() {
			BeforeEach(func() {
				clientFake.LocationReturns(client.Location{}, fmt.Errorf("something went wrong"))
			})

			It("returns error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
