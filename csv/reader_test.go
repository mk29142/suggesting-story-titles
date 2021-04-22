package csv_test

import (
	"github.com/mk29142/suggesting-story-titles/csv"
	"github.com/mk29142/suggesting-story-titles/domain"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Reader", func() {
   var (
   	reader csv.Reader
   	file string

   	res []domain.Metadata
   	err error
	 )

   BeforeEach(func() {
   	file = "./fixtures/test.csv"
   	reader = csv.NewReader()
	 })

   JustBeforeEach(func() {
   	res, err = reader.Read(file)
	 })

   When("success", func() {
   	It("should parse file", func() {
   		format := "2006-01-02 15:04:05"
   		timestamp1, _ := time.Parse(format, "2020-03-30 14:12:19")
   		timestamp2, _ := time.Parse(format, "2020-03-30 14:20:10")
   		timestamp3, _ := time.Parse(format, "2020-03-30 14:32:02")

   		expected := []domain.Metadata{
				{
					Timestamp: timestamp1,
					Latitude:  40.728808,
					Longitude: -73.996106,
				},
				{
					Timestamp: timestamp2,
					Latitude:  40.728656,
					Longitude: -73.998790,
				},
				{
					Timestamp: timestamp3,
					Latitude:  40.727160,
					Longitude: -73.996044,
				},
			}

			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(Equal(expected))
		})
	 })

})
