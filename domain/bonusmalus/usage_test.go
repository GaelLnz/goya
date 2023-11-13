package bonusmalus_test

import (
	"testing"

	"github.com/GaelLnz/goya/domain/bonusmalus"
	"github.com/stretchr/testify/assert"
)

func TestUsageValid(t *testing.T) {
	assert.True(t, bonusmalus.BonusMalusPrivateUsage.IsValid())
	assert.True(t, bonusmalus.BonusMalusProfessionnalUsage.IsValid())
}

func TestUsageInvalid(t *testing.T) {
	assert.False(t, bonusmalus.BonusMalusUsage("invalid").IsValid())
}
