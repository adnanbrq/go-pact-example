package main

import (
	"fmt"
	"testing"

	"github.com/pact-foundation/pact-go/v2/consumer"
	"github.com/stretchr/testify/assert"
)

func TestConsumer(t *testing.T) {
	mockProvider, err := consumer.NewV4Pact(consumer.MockHTTPProviderConfig{
		Consumer: "Jokes-CLI",
		Provider: "Jokes-API",
		Host:     "127.0.0.1",
	})

	assert.NoError(t, err)

	err = mockProvider.
		AddInteraction().
		Given("a list of jokes").
		UponReceiving("a request for a random joke").
		WithRequest("GET", "/").
		WillRespondWith(200, func(b *consumer.V4ResponseBuilder) {
			b.BodyMatch(Joke{})
		}).
		ExecuteTest(t, func(msc consumer.MockServerConfig) error {
			url := fmt.Sprintf("http://%s:%d", msc.Host, msc.Port)
			_, err := getRandomJoke(url)

			assert.NoError(t, err)
			return err
		})

	assert.NoError(t, err)
}
