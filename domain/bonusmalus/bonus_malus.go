package bonusmalus

import (
	"fmt"
)

type BonusMalus struct {
	drivingYears uint
	accidents    uint
	usage        BonusMalusUsage
}

func NewBonusMalus(drivingYears, accidents uint, usage BonusMalusUsage) (*BonusMalus, error) {
	if !usage.IsValid() {
		return nil, fmt.Errorf("%s is not a valid usage", usage)
	}

	return &BonusMalus{
		drivingYears: normalizeDrivingYears(drivingYears),
		accidents:    normalizeAccidents(accidents),
		usage:        usage,
	}, nil
}

func (b BonusMalus) Score() int {
	const accidentPenalty int = 5

	score := b.initialScore() - int(b.drivingYears) + int(b.accidents)*accidentPenalty

	return b.normalizeScore(score)
}

func (b BonusMalus) initialScore() int {
	if b.usage == BonusMalusPrivateUsage {
		return 11
	}

	return 14
}

func (b BonusMalus) normalizeScore(score int) int {
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
func (b BonusMalus) minScore() int {
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
