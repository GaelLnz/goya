package graphql_test

import (
	"context"
	"testing"

	"github.com/GaelLnz/goya/api/graphql"
	"github.com/GaelLnz/goya/api/graphql/model"
	"github.com/GaelLnz/goya/api/graphql/testutils"
	"github.com/stretchr/testify/assert"
)

func TestBonusMalusScorePrivate(t *testing.T) {
	testutils.WithTestResolver(func(resolver *graphql.Resolver) {
		score, err := resolver.Query().BonusMalusScore(context.Background(), 1, 2, model.BonusMalusUsagePrivate)
		assert.NoError(t, err)
		assert.Equal(t, 20, score)
	})
}

func TestBonusMalusScoreProfesionnal(t *testing.T) {
	testutils.WithTestResolver(func(resolver *graphql.Resolver) {
		score, err := resolver.Query().BonusMalusScore(context.Background(), 1, 2, model.BonusMalusUsageProfesionnal)
		assert.NoError(t, err)
		assert.Equal(t, 22, score)
	})
}

func TestBonusMalusScoreNegativeDrivingYears(t *testing.T) {
	testutils.WithTestResolver(func(resolver *graphql.Resolver) {
		score, err := resolver.Query().BonusMalusScore(context.Background(), -1, 0, model.BonusMalusUsageProfesionnal)
		assert.Error(t, err)
		assert.Equal(t, 0, score)
	})
}

func TestBonusMalusScoreNegativeAccidents(t *testing.T) {
	testutils.WithTestResolver(func(resolver *graphql.Resolver) {
		score, err := resolver.Query().BonusMalusScore(context.Background(), 0, -1, model.BonusMalusUsageProfesionnal)
		assert.Error(t, err)
		assert.Equal(t, 0, score)
	})
}

func TestBonusMalusScoreInvalidUsage(t *testing.T) {
	testutils.WithTestResolver(func(resolver *graphql.Resolver) {
		score, err := resolver.Query().BonusMalusScore(context.Background(), 0, 0, "invalid")
		assert.Error(t, err)
		assert.Equal(t, 0, score)
	})
}
