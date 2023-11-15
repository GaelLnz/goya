package bonusmalus

import (
	"fmt"
)

type bonusMalus struct {
	drivingYears uint
	accidents    uint
	usage        BonusMalusUsage
}

func NewBonusMalus(drivingYears, accidents uint, usage BonusMalusUsage) (*bonusMalus, error) {
	if !usage.IsValid() {
		return nil, fmt.Errorf("%s is not a valid usage", usage)
	}

	return &bonusMalus{
		drivingYears: normalizeDrivingYears(drivingYears),
		accidents:    normalizeAccidents(accidents),
		usage:        usage,
	}, nil
}

func (b bonusMalus) Score() int {
	const accidentPenalty int = 5

	score := b.initialScore() - int(b.drivingYears) + int(b.accidents)*accidentPenalty

	return b.normalizeScore(score)
}

func (b bonusMalus) initialScore() int {
	if b.usage == BonusMalusPrivateUsage {
		return 11
	}

	return 14
}

func (b bonusMalus) normalizeScore(score int) int {
	const maxScore int = 22
	if score > maxScore {
		return maxScore
	}

	minScore := b.minScore()
	if score < minScore {
		return minScore
	}

	return score
}
func (b bonusMalus) minScore() int {
	if b.accidents < 2 {
		return -2
	}

	return 3
}

func normalizeDrivingYears(drivingYears uint) uint {
	return min(drivingYears, 22)
}

func normalizeAccidents(accidents uint) uint {
	return min(accidents, 2)
}
