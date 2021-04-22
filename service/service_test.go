package service_test

import (
	"github.com/mk29142/suggesting-story-titles/domain"
	"github.com/mk29142/suggesting-story-titles/generator"
	"github.com/mk29142/suggesting-story-titles/service"
	"github.com/mk29142/suggesting-story-titles/service/internal"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Suggestor", func() {
  var (
  	fakeGenerator1 *internal.FakeGenerator
  	fakeGenerator2 *internal.FakeGenerator
  	fakeGenerator3 *internal.FakeGenerator

  	locations []domain.Location

  	s service.Service

  	res []service.Suggestion
  	title1 string
  	title2 string
  	title3 string
	)

  BeforeEach(func() {
  	fakeGenerator1 = new(internal.FakeGenerator)
  	fakeGenerator2 = new(internal.FakeGenerator)
  	fakeGenerator3 = new(internal.FakeGenerator)

  	fakeGenerator1.GenerateReturns(generator.Suggestion{Title: title1})
  	fakeGenerator2.GenerateReturns(generator.Suggestion{Title: title2})
  	fakeGenerator3.GenerateReturns(generator.Suggestion{Title: title3})

  	s = service.NewSuggestor(fakeGenerator1, fakeGenerator2, fakeGenerator3)

  	locations = []domain.Location{
  		{
				Latitude:  50,
				Longitude: -40,
				Timestamp: time.Time{},
				Name:      "name1",
		 },
			{
				Latitude:  50,
				Longitude: -40,
				Timestamp: time.Time{},
				Name:      "name2",
			},
			{
				Latitude:  50,
				Longitude: -40,
				Timestamp: time.Time{},
				Name:      "name2",
			},
  	}
	})

  JustBeforeEach(func() {
		res = s.Suggestions(locations)
	})

  When("Success", func() {
		It("uses generators to give suggestions", func() {

			expected := []string {title1, title2, title3}

			for _, suggestion := range res {
				title := suggestion.Title
				Expect(stringInSlice(title, expected)).To(BeTrue())
			}
		})
	})
})

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
