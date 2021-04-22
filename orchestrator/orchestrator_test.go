package orchestrator_test

import (
	"github.com/mk29142/suggesting-story-titles/client"
	"github.com/mk29142/suggesting-story-titles/domain"
	"github.com/mk29142/suggesting-story-titles/orchestrator"
	"github.com/mk29142/suggesting-story-titles/orchestrator/internal"
	"github.com/mk29142/suggesting-story-titles/service"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Orchestrator", func() {
   var (
   		fakeSuggestor *internal.FakeSuggestor
   		fakeClient    *internal.FakeClient

   		o orchestrator.Orchestrator

   		data []domain.Metadata

   		res []domain.Title
	 )

   BeforeEach(func() {
   	fakeSuggestor = new(internal.FakeSuggestor)
   	fakeClient = new(internal.FakeClient)

   	o = orchestrator.New(fakeSuggestor, fakeClient)

   	data = []domain.Metadata{
			{
				Timestamp: time.Time{},
				Latitude:  10,
				Longitude: -40,
			},
			{
				Timestamp: time.Time{},
				Latitude:  10,
				Longitude: -40,
			},
			{
				Timestamp: time.Time{},
				Latitude:  10,
				Longitude: -40,
			},
		}
	 })

   JustBeforeEach(func() {
   	res = o.Titles(data)
	 })


   When("success", func() {
   		BeforeEach(func() {
   			fakeClient.LocationReturns(client.Location{
					Latitude:  10,
					Longitude: -40,
					Timestamp: time.Time{},
					Name:      "London",
				}, nil)

   			fakeSuggestor.SuggestionsReturns([]service.Suggestion{
					{
						Latitude:  10,
						Longitude: -40,
						Timestamp: time.Time{},
						Name:      "London",
						Title:     "woop woop london",
					},
					{
						Latitude:  10,
						Longitude: -40,
						Timestamp: time.Time{},
						Name:      "London",
						Title:     "woop woop london",
					},
					{
						Latitude:  10,
						Longitude: -40,
						Timestamp: time.Time{},
						Name:      "London",
						Title:     "woop woop london",
					},
				})
			})

   		It("gets all titles", func() {
   			By("calling client", func() {
					Expect(fakeClient.LocationCallCount()).To(Equal(3))
				})

				By("calling the suggestor", func() {
					Expect(fakeSuggestor.SuggestionsCallCount()).To(Equal(1))
				})

   			By("correctly parsing location", func() {
					expected := []domain.Location{
						{
							Latitude:  10,
							Longitude: -40,
							Timestamp: time.Time{},
							Name:      "London",
						},
						{
							Latitude:  10,
							Longitude: -40,
							Timestamp: time.Time{},
							Name:      "London",
						},
						{
							Latitude:  10,
							Longitude: -40,
							Timestamp: time.Time{},
							Name:      "London",
						},
					}

					Expect(fakeSuggestor.SuggestionsArgsForCall(0)).To(Equal(expected))
				})

				By("correctly suggestions suggestions", func() {
					expected := []domain.Title{
						{
							Latitude:  10,
							Longitude: -40,
							Timestamp: time.Time{},
							Name:      "London",
							Title:     "woop woop london",
						},
						{
							Latitude:  10,
							Longitude: -40,
							Timestamp: time.Time{},
							Name:      "London",
							Title:     "woop woop london",
						},
						{
							Latitude:  10,
							Longitude: -40,
							Timestamp: time.Time{},
							Name:      "London",
							Title:     "woop woop london",
						},
					}

					Expect(res).To(Equal(expected))
				})
			})
	 })
})
