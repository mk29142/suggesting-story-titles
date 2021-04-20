package generator_test

import (
	"github.com/mk29142/suggesting-story-titles/domain"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"

	"github.com/mk29142/suggesting-story-titles/generator"
)

var _ = Describe("GenericGenerator", func() {
	When("Generic", func() {
		var (
			format = "2006-01-02 15:04:05"
			location domain.Location
			timestamp time.Time

			g generator.GenericGenerator
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

			g = generator.NewGenericGenerator()
		})

		It("returns generic titles", func() {
			got := g.Generate(location)

			expected := []string {"A fun trip to London", "Fun times in London", "Take me back to London", "Time well spent in London"}

			Expect(stringInSlice(got.Title, expected)).To(BeTrue())
		})
	})
})
