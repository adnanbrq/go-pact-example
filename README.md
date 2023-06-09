# Pact Example in GoLang

This is an exmample of how to use pact in GoLang.\
This example focuses on how a Contract can be Created and Verified. Publishing of contracts is not covered.

## Why Pact

Pact can be used to ensure that a Contract like a REST-API is served and consumable like expected. \
We have two services in this example project.\
**Jokes-API** is a HTTP Service that returns a random joke out of a given list.\
**Jokes-CLI** is a CLI Application that sends a Request to the _Jokes-API_ Service and renders the
retrieved joke in the Console.\

In Pact we have a Consumer that describes how it consumes a Provider. We use Pact and it's DSL to generate a Contract that later can be used by the consumed Provider to see if it matches the requirements of it's consumers. \

### But why though?

Through Contract lets us know when we breake a Contract between Services. If a Provider does not fullfil it's consumers requirements means that we would need to stop a deployment as it would possibly break our System.

> **Warning**\
> Pact is not a integration or unit testing tool.
