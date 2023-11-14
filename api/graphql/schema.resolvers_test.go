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
		score, err := resolver.Query().BonusMalusScore(context.Background(), 0, 0, model.BonusMalusUsagePrivate)
		assert.NoError(t, err)
		assert.Equal(t, 11, score)
	})
}

func TestBonusMalusScoreProfesionnal(t *testing.T) {
	testutils.WithTestResolver(func(resolver *graphql.Resolver) {
		score, err := resolver.Query().BonusMalusScore(context.Background(), 0, 0, model.BonusMalusUsageProfessional)
		assert.NoError(t, err)
		assert.Equal(t, 14, score)
	})
}
