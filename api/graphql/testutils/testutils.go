package testutils

import (
	"log"

	"github.com/GaelLnz/goya/api/graphql"
)

func WithTestResolver(testFunction func(resolver *graphql.Resolver)) {
	resolver := graphql.NewResolver(log.Default())

	testFunction(&resolver)
}
