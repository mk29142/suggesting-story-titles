package generator_test

import (
	"github.com/mk29142/suggesting-story-titles/domain"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"

	"github.com/mk29142/suggesting-story-titles/generator"
)

var _ = Describe("Seasonbasedgenerator", func() {
	When("December", func() {
		var (
			format = "2006-01-02 15:04:05"
			location domain.Location
			timestamp time.Time

			g generator.SeasonBasedGenerator
		)

		BeforeEach(func() {
			t, err := time.Parse(format, "2019-12-30 07:12:19")
			Expect(err).NotTo(HaveOccurred())

			timestamp = t

			location = domain.Location{
				Latitude:  40.703717,
				Longitude: -74.016094,
				Timestamp: timestamp,
				Name:      "London",
			}

			g = generator.NewSeasonBasedGenerator()
		})

		It("returns messages relating to winter", func() {
			got := g.Generate(location)

			expected := []string {"A chilly winter in London", "London in the winter"}

			Expect(stringInSlice(got.Title, expected)).To(BeTrue())
		})
	})

	When("February", func() {
		var (
			format = "2006-01-02 15:04:05"
			location domain.Location
			timestamp time.Time

			g generator.SeasonBasedGenerator
		)

		BeforeEach(func() {
			t, err := time.Parse(format, "2019-02-28 07:12:19")
			Expect(err).NotTo(HaveOccurred())

			timestamp = t

			location = domain.Location{
				Latitude:  40.703717,
				Longitude: -74.016094,
				Timestamp: timestamp,
				Name:      "London",
			}

			g = generator.NewSeasonBasedGenerator()
		})

		It("returns messages relating to winter", func() {
			got := g.Generate(location)

			expected := []string {"A chilly winter in London", "London in the winter"}

			Expect(stringInSlice(got.Title, expected)).To(BeTrue())
		})
	})

	When("April", func() {
		var (
			format = "2006-01-02 15:04:05"
			location domain.Location
			timestamp time.Time

			g generator.SeasonBasedGenerator
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

			g = generator.NewSeasonBasedGenerator()
		})

		It("returns messages relating to spring", func() {
			got := g.Generate(location)

			expected := []string {"A blooming spring in London", "London in the spring"}

			Expect(stringInSlice(got.Title, expected)).To(BeTrue())
		})
	})

	When("July", func() {
		var (
			format = "2006-01-02 15:04:05"
			location domain.Location
			timestamp time.Time

			g generator.SeasonBasedGenerator
		)

		BeforeEach(func() {
			t, err := time.Parse(format, "2019-06-30 07:12:19")
			Expect(err).NotTo(HaveOccurred())

			timestamp = t

			location = domain.Location{
				Latitude:  40.703717,
				Longitude: -74.016094,
				Timestamp: timestamp,
				Name:      "London",
			}

			g = generator.NewSeasonBasedGenerator()
		})

		It("returns messages relating to summer", func() {
			got := g.Generate(location)

			expected := []string {"A beautiful summer in London", "London in the summer"}

			Expect(stringInSlice(got.Title, expected)).To(BeTrue())
		})
	})

	When("November", func() {
		var (
			format = "2006-01-02 15:04:05"
			location domain.Location
			timestamp time.Time

			g generator.SeasonBasedGenerator
		)

		BeforeEach(func() {
			t, err := time.Parse(format, "2019-11-30 07:12:19")
			Expect(err).NotTo(HaveOccurred())

			timestamp = t

			location = domain.Location{
				Latitude:  40.703717,
				Longitude: -74.016094,
				Timestamp: timestamp,
				Name:      "London",
			}

			g = generator.NewSeasonBasedGenerator()
		})

		It("returns messages relating to autumn", func() {
			got := g.Generate(location)

			expected := []string {"A picturesque autumn in London", "London in the autumn"}

			Expect(stringInSlice(got.Title, expected)).To(BeTrue())
		})
	})
})
