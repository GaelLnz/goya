package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"errors"

	"github.com/GaelLnz/goya/api/graphql/model"
	"github.com/GaelLnz/goya/domain/bonusmalus"
)

// BonusMalusScore is the resolver for the bonusMalusScore field.
func (r *queryResolver) BonusMalusScore(ctx context.Context, drivingYears int, accidents int, usage model.BonusMalusUsage) (int, error) {
	r.logger.Printf("querying bonus-malus for %d driving years, %d accidents and %s usage", drivingYears, accidents, usage.String())

	// Would've been great to use uint as parameters here, but for some reason gqlgen seems to do a poor job at handling them.
	// Negative numbers passed to uint parameters do NOT result in an API error and are badly casted. I have no time to investigate further. More here https://gqlgen.com/reference/scalars/
	if drivingYears < 0 {
		return 0, errors.New("driving years must be a positive integer")
	}

	if accidents < 0 {
		return 0, errors.New("accidents must be a positive integer")
	}

	bonusMalus, err := bonusmalus.NewBonusMalus(uint(drivingYears), uint(accidents), bonusmalus.BonusMalusUsage(usage))
	if err != nil {
		return 0, err
	}

	return bonusMalus.Score(), nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
