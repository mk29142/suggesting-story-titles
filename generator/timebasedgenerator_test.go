package generator_test

import (
	"github.com/mk29142/suggesting-story-titles/domain"
	 "github.com/mk29142/suggesting-story-titles/generator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Timebasedgenerator", func() {

	Describe("Morning", func() {
		When("7am", func() {
			var (
				format = "2006-01-02 15:04:05"
				location domain.Location
				timestamp time.Time

				g generator.TimeBasedGenerator
			)

			BeforeEach(func() {
				t, err := time.Parse(format, "2019-03-30 07:12:19")
				Expect(err).NotTo(HaveOccurred())

				timestamp = t

				location = domain.Location{
					Latitude:  40.703717,
					Longitude: -74.016094,
					Timestamp: timestamp,
					Name:      "London",
				}

				g = generator.NewTimeBasedGenerator()
			})

			It("returns messages relating to the morning", func() {
				got := g.Generate(location)

				expected := []string {"A lovely morning in London", "A peaceful morning in London", "London in the morning"}

				Expect(stringInSlice(got.Title, expected)).To(BeTrue())
			})
		})
		When("midnight", func() {
			var (
				format = "2006-01-02 15:04:05"
				location domain.Location
				timestamp time.Time

				g generator.TimeBasedGenerator
			)

			BeforeEach(func() {
				t, err := time.Parse(format, "2019-03-30 00:00:00")
				Expect(err).NotTo(HaveOccurred())

				timestamp = t

				location = domain.Location{
					Latitude:  40.703717,
					Longitude: -74.016094,
					Timestamp: timestamp,
					Name:      "London",
				}

				g = generator.NewTimeBasedGenerator()
			})

			It("returns messages relating to the morning", func() {
				got := g.Generate(location)

				expected := []string {"A lovely morning in London", "A peaceful morning in London", "London in the morning"}

				Expect(stringInSlice(got.Title, expected)).To(BeTrue())
			})
		})
	})

	Describe("Afternoon", func() {
		When("3pm", func() {
			var (
				format = "2006-01-02 15:04:05"
				location domain.Location
				timestamp time.Time

				g generator.TimeBasedGenerator
			)

			BeforeEach(func() {
				t, err := time.Parse(format, "2019-03-30 15:12:19")
				Expect(err).NotTo(HaveOccurred())

				timestamp = t

				location = domain.Location{
					Latitude:  40.703717,
					Longitude: -74.016094,
					Timestamp: timestamp,
					Name:      "London",
				}

				g = generator.NewTimeBasedGenerator()
			})

			It("returns messages relating to the afternoon", func() {
				got := g.Generate(location)

				expected := []string {"Enjoying the afternoon in London", "A fun afternoon in London", "London in the afternoon"}

				Expect(stringInSlice(got.Title, expected)).To(BeTrue())
			})
		})
		When("noon", func() {
			var (
				format = "2006-01-02 15:04:05"
				location domain.Location
				timestamp time.Time

				g generator.TimeBasedGenerator
			)

			BeforeEach(func() {
				t, err := time.Parse(format, "2019-03-30 12:00:00")
				Expect(err).NotTo(HaveOccurred())

				timestamp = t

				location = domain.Location{
					Latitude:  40.703717,
					Longitude: -74.016094,
					Timestamp: timestamp,
					Name:      "London",
				}

				g = generator.NewTimeBasedGenerator()
			})

			It("returns messages relating to the afternoon", func() {
				got := g.Generate(location)

				expected := []string {"Enjoying the afternoon in London", "A fun afternoon in London", "London in the afternoon"}

				Expect(stringInSlice(got.Title, expected)).To(BeTrue())
			})
		})
	})

	Describe("Evening", func() {
		When("8pm", func() {
			var (
				format = "2006-01-02 15:04:05"
				location domain.Location
				timestamp time.Time

				g generator.TimeBasedGenerator
			)

			BeforeEach(func() {
				t, err := time.Parse(format, "2019-03-30 20:12:19")
				Expect(err).NotTo(HaveOccurred())

				timestamp = t

				location = domain.Location{
					Latitude:  40.703717,
					Longitude: -74.016094,
					Timestamp: timestamp,
					Name:      "London",
				}

				g = generator.NewTimeBasedGenerator()
			})

			It("returns messages relating to the evening", func() {
				got := g.Generate(location)

				expected := []string {"Evening in London", "A quiet evening walk in London", "A blissful night spent in London"}

				Expect(stringInSlice(got.Title, expected)).To(BeTrue())
			})
		})

		When("just before midnight", func() {
			var (
				format = "2006-01-02 15:04:05"
				location domain.Location
				timestamp time.Time

				g generator.TimeBasedGenerator
			)

			BeforeEach(func() {
				t, err := time.Parse(format, "2019-03-30 23:59:00")
				Expect(err).NotTo(HaveOccurred())

				timestamp = t

				location = domain.Location{
					Latitude:  40.703717,
					Longitude: -74.016094,
					Timestamp: timestamp,
					Name:      "London",
				}

				g = generator.NewTimeBasedGenerator()
			})

			It("returns messages relating to the evening", func() {
				got := g.Generate(location)

				expected := []string {"Evening in London", "A quiet evening walk in London", "A blissful night spent in London"}

				Expect(stringInSlice(got.Title, expected)).To(BeTrue())
			})
		})
	})
})

