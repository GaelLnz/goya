package testutils

import "github.com/GaelLnz/goya/api/graphql"

func WithTestResolver(testFunction func(resolver *graphql.Resolver)) {
	resolver := &graphql.Resolver{}

	testFunction(resolver)
}
