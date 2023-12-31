# Technical test for Goya

## Overview

The goal of this challenge is to create an API that can be used to compute the Belgium driver's bonus-malus score.
This score depends on the number of driving years, the number of accidents and the driving usage. The score is in the range -2 (best) to 22 (worst).

## Architectural & technical choices

### Project structure

The project has the following structure:
* `domain/` => contains the business logic
* `api/graphql/` => contains the graphql api
* `cmd/server/` => contains the entry point of this application
* `Dockerfile` => a multi-stage build dockerfile used to build & run the application
* `Makefile` => a makefile to simplify developpement

### Domain Driven Design

I chose to use the DDD pattern as it focuses on creating a shared understanding of the problem domain between technical and non-technical team members. This approach also facilitates the modularity and readability of the code.

### GraphQL

GraphQL provides a powerful and flexible query language for fetching data. It has many advantages but it is definitely overkill for such small usecase. The only real reason to implement a GraphQL API here is pure fun.

#### gqlgen

A simple & powerful library for building GraphQL servers. More at [https://gqlgen.com/](https://gqlgen.com/).

To update the GraphQL schema, you will need to modify the `api/graphql/schema.graphqls` file. It is then necessary to re-generate code associated to this schema using:

> `make generate`

## Try it yourself

### Dependencies

#### Run

To run the application, you will need to install on your machine:
- `docker`
- `make`

#### Test & lint

To test and lint, you will need to install on your machine:
- `go >= 1.21`
- `golangci-lint`

### Running the application

The easiest way to run this application is via `docker`.

> `make run`

#### Usage

This application exposes a GraphQL API at `http://localhost:8080/query` by default. You can manually interact with the API using the built-in GraphQL playground at `http://localhost:8080/`.

There is only one query available, `bonusMalusScore` which returns the bonus-malus score as an integer. Further user documentation can be retrieved from the GraphQL schema accessible on the playground.

### Testing the application

You can test the application using:

> `make test`

### Lint

You can lint the code using:

> `make lint`

## Possible enhancements

- [ ] Improve this README
- [ ] Run tests in CI as a github action
- [ ] Convert `bonusMalusScore` query parameters to uint
