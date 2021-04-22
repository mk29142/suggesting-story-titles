package client_test

import (
	"bytes"
	"fmt"
	"github.com/mk29142/suggesting-story-titles/client"
	"github.com/mk29142/suggesting-story-titles/domain"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"

	"github.com/mk29142/suggesting-story-titles/client/internal"
)

var _ = Describe("Client", func() {
	var (
		apiToken   string
		clientFake *internal.FakeClient
		service    client.Client
	)

	BeforeEach(func() {
		clientFake = new(internal.FakeClient)
		clientFake.DoReturns(&http.Response{
			Body:       ioutil.NopCloser(new(bytes.Buffer)),
			StatusCode: http.StatusOK,
		}, nil)
		apiToken = "super-secret-header-value"
		service = client.New(apiToken, clientFake)
	})

	Describe("Name", func() {
		var (
			input domain.Metadata
			result client.Location
			err error
		)

		JustBeforeEach(func() {
			result, err = service.Location(input)
		})

		When("given valid coordinates", func() {
			BeforeEach(func() {
				input = domain.Metadata{
					Latitude:  50.123,
					Longitude: 0.4021,
				}
				resp := `{
					"features": [
              { 
                "text": "London"
  						}
						]
				}`
				clientFake.DoReturns(&http.Response{
					Body:       ioutil.NopCloser(bytes.NewBufferString(resp)),
					StatusCode: http.StatusOK,
				}, nil)
			})


			It("should not error", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("should decode response correctly", func() {
				Expect(result).To(Equal(client.Location{
					Latitude:  input.Latitude,
					Longitude: input.Longitude,
					Name:      "London",
				}))
			})

			It("should call the client correctly", func() {
				Expect(clientFake.DoCallCount()).To(Equal(1))
				got := clientFake.DoArgsForCall(0)

				By("performing a get request", func() {
					Expect(got.Method).To(Equal(http.MethodGet))
				})

				By("using the correct url", func() {
					expectedUrl :=
						fmt.Sprintf("https://api.mapbox.com/geocoding/v5/mapbox.places/%f,%f.json?types=place&limit=1&access_token=%s",
							input.Longitude, input.Latitude, apiToken)

					Expect(got.URL.String()).To(Equal(expectedUrl))
				})
			})
		})

		When("response cannot be decoded", func() {
			BeforeEach(func() {
				input = domain.Metadata{
					Latitude:  50.123,
					Longitude: 0.4021,
				}
				resp := `not json`
				clientFake.DoReturns(&http.Response{
					Body:       ioutil.NopCloser(bytes.NewBufferString(resp)),
					StatusCode: http.StatusOK,
				}, nil)
			})

			It("should error", func() {
				Expect(err).To(HaveOccurred())
			})
		})

		When("request does not return 200", func() {
			BeforeEach(func() {
				input = domain.Metadata{
					Latitude:  50.123,
					Longitude: 0.4021,
				}
				resp := `{}`
				clientFake.DoReturns(&http.Response{
					Body:       ioutil.NopCloser(bytes.NewBufferString(resp)),
					StatusCode: http.StatusNotFound,
				}, nil)
			})

			It("should return error", func() {
				Expect(clientFake.DoCallCount()).To(Equal(1))
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("unexpected status code"))
			})
		})

		When("request to api fails", func() {
			BeforeEach(func() {
				input = domain.Metadata{
					Latitude:  50.123,
					Longitude: 0.4021,
				}
				clientFake.DoReturns(&http.Response{}, fmt.Errorf("something went wrong"))
			})

			It("should return error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
