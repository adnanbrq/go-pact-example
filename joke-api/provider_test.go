package main

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/pact-foundation/pact-go/v2/log"
	"github.com/pact-foundation/pact-go/v2/provider"
	"github.com/pact-foundation/pact-go/v2/version"
	"github.com/stretchr/testify/assert"
)

func TestProvider(t *testing.T) {
	// Optional
	log.SetLogLevel("TRACE")
	version.CheckVersion()

	// Start the server in the background
	go startHttpServer("127.0.0.1:3031")

	go func() {
		err := provider.
			NewVerifier().
			VerifyProvider(t, provider.VerifyRequest{
				ProviderBaseURL: "http://127.0.0.1:3031",
				Provider:        "Jokes-API",
				PactFiles: []string{
					filepath.ToSlash(fmt.Sprintf("%s/Jokes-CLI-Jokes-API.json", "../joke-cli/pacts")),
				},
			})

		assert.NoError(t, err)
	}()
}
