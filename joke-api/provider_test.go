package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/adnanbrq/goenv"
	"github.com/pact-foundation/pact-go/v2/log"
	"github.com/pact-foundation/pact-go/v2/provider"
	"github.com/pact-foundation/pact-go/v2/version"
	"github.com/stretchr/testify/assert"
)

func TestProvider(t *testing.T) {
	// Load .env variables into os.Getenv
	goenv.Load(false)

	// Optional
	log.SetLogLevel("TRACE")
	version.CheckVersion()

	// Start the server in the background
	go startHttpServer("127.0.0.1:3031")

	err := provider.
		NewVerifier().
		VerifyProvider(t, provider.VerifyRequest{
			ProviderBaseURL: "http://127.0.0.1:3031",
			ProviderVersion: "1.0.0",
			Provider:        "Jokes-API",
			PactFiles: []string{
				filepath.ToSlash(fmt.Sprintf("%s/Jokes-CLI-Jokes-API.json", "../joke-cli/pacts")),
			},
			BrokerURL:                  os.Getenv("PACT_BROKER_URL"),
			BrokerUsername:             os.Getenv("PACT_BROKER_USERNAME"),
			BrokerPassword:             os.Getenv("PACT_BROKER_PASSWORD"),
			PublishVerificationResults: true,
		})

	assert.NoError(t, err)
}
