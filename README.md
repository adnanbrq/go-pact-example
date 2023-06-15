# Pact Example in GoLang

This is an exmample of how to use pact in GoLang.\
This example focuses on how a Contract can be Created and Verified.\
\
As far as I know, publishing a contract is not yet supported by the API and can only be done through a direct http request to the Pact-Broker or through a Pact-CLI\
_Publishing Verification results is supported though as seen in **provider_test**_\
\
Uploading a Contract to the Broker could be done using Postman

> URL: http://HOST:PORT/pacts/provider/Jokes-API/consumer/Jokes-CLI/version/$PROVIDER-VERSION\
> Header: Basic $USER:$PASS\
> Body: Contract JSON\
> Method: PUT

## Why Pact

Pact can be used to ensure that a Contract like a REST-API is served and consumable like expected. \
We have two services in this example project.\
**Jokes-API** is a HTTP Service that returns a random joke out of a given list.\
**Jokes-CLI** is a CLI Application that sends a Request to the _Jokes-API_ Service and renders the retrieved joke in the Console.

In Pact we have a Consumer that describes how it consumes a Provider. We use Pact and it's DSL to generate a Contract that later can be used by the consumed Provider to see if it matches the requirements of it's consumers.

### But why though?

Contract Testing with for example Pact can have some benefits. One being that we can get notified if one Provider does not fullfil one of it's Consumers requirements.\
If done in a pipeline, we can stop a rollout / deployment and notify the developing team when we know that a Provider Verification failed. This prevents from breaking our system.

> **Warning**\
> Pact is not a integration or unit testing tool.

## Dependencies

- [github.com/pact-foundation/pact-go](github.com/pact-foundation/pact-go/v2)
  API to create and verify contracts
- [github.com/stretchr/testify](github.com/stretchr/testify)
  Used for assertions on tests
- [github.com/adnanbrq/goenv](github.com/adnanbrq/goenv)
  This will parse a .env file and make the values available via os.Getenv
