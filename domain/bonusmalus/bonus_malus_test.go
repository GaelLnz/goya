package bonusmalus_test

import (
	"fmt"
	"testing"

	"github.com/GaelLnz/goya/domain/bonusmalus"
	"github.com/stretchr/testify/assert"
)

type scoreTestCase struct {
	drivingYears uint
	expected     int
}

func TestScorePrivate(t *testing.T) {
	testCases0Accident := []scoreTestCase{
		{0, 11},
		{1, 10},
		{2, 9},
		{3, 8},
		{4, 7},
		{5, 6},
		{6, 5},
		{7, 4},
		{8, 3},
		{9, 2},
		{10, 1},
		{11, 0},
		{12, -1},
		{13, -2},
		{14, -2},
		{15, -2},
		{16, -2},
		{17, -2},
		{18, -2},
		{19, -2},
		{20, -2},
		{21, -2},
		{22, -2},
		{23, -2},
		{99, -2},
	}
	testCases1Accident := []scoreTestCase{
		{0, 16},
		{1, 15},
		{2, 14},
		{3, 13},
		{4, 12},
		{5, 11},
		{6, 10},
		{7, 9},
		{8, 8},
		{9, 8},
		{10, 6},
		{11, 5},
		{12, 4},
		{13, 3},
		{14, 2},
		{15, 1},
		{16, 0},
		{17, -1},
		{18, -2},
		{19, -2},
		{20, -2},
		{21, -2},
		{22, -2},
		{23, -2},
		{99, -2},
	}
	testCases2Accident := []scoreTestCase{
		{0, 21},
		{1, 20},
		{2, 19},
		{3, 18},
		{4, 17},
		{5, 16},
		{6, 15},
		{7, 14},
		{8, 13},
		{9, 12},
		{10, 11},
		{11, 10},
		{12, 9},
		{13, 8},
		{14, 7},
		{15, 6},
		{16, 5},
		{17, 4},
		{18, 3},
		{19, 3},
		{20, 3},
		{21, 3},
		{22, 3},
		{23, 3},
		{99, 3},
	}

	assertScoreTable(t, testCases0Accident, 0 /* accident */, bonusmalus.BonusMalusPrivateUsage)
	assertScoreTable(t, testCases1Accident, 1 /* accident */, bonusmalus.BonusMalusPrivateUsage)
	assertScoreTable(t, testCases2Accident, 2 /* accident */, bonusmalus.BonusMalusPrivateUsage)
	assertScoreTable(t, testCases2Accident, 3 /* accident */, bonusmalus.BonusMalusPrivateUsage)
	assertScoreTable(t, testCases2Accident, 99 /* accident */, bonusmalus.BonusMalusPrivateUsage)
}

func TestScorePublic(t *testing.T) {
	testCases0Accident := []scoreTestCase{
		{0, 14},
		{1, 13},
		{2, 12},
		{3, 11},
		{4, 10},
		{5, 9},
		{6, 8},
		{7, 7},
		{8, 6},
		{9, 5},
		{10, 4},
		{11, 3},
		{12, 2},
		{13, 1},
		{14, 0},
		{15, -1},
		{16, -2},
		{17, -2},
		{18, -2},
		{19, -2},
		{20, -2},
		{21, -2},
		{22, -2},
		{23, -2},
		{99, -2},
	}
	testCases1Accident := []scoreTestCase{
		{0, 19},
		{1, 18},
		{2, 17},
		{3, 16},
		{4, 15},
		{5, 14},
		{6, 13},
		{7, 12},
		{8, 11},
		{9, 10},
		{10, 9},
		{11, 8},
		{12, 7},
		{13, 6},
		{14, 5},
		{15, 4},
		{16, 3},
		{17, 2},
		{18, 1},
		{19, 0},
		{20, -1},
		{21, -2},
		{22, -2},
		{23, -2},
		{99, -2},
	}
	testCases2Accident := []scoreTestCase{
		{0, 22},
		{1, 22},
		{2, 22},
		{3, 21},
		{4, 20},
		{5, 19},
		{6, 18},
		{7, 17},
		{8, 16},
		{9, 15},
		{10, 14},
		{11, 13},
		{12, 12},
		{13, 11},
		{14, 10},
		{15, 9},
		{16, 8},
		{17, 7},
		{18, 6},
		{19, 5},
		{20, 4},
		{21, 3},
		{22, 3},
		{23, 3},
		{99, 3},
	}

	assertScoreTable(t, testCases0Accident, 0 /* accident */, bonusmalus.BonusMalusPublicUsage)
	assertScoreTable(t, testCases1Accident, 1 /* accident */, bonusmalus.BonusMalusPublicUsage)
	assertScoreTable(t, testCases2Accident, 2 /* accident */, bonusmalus.BonusMalusPublicUsage)
	assertScoreTable(t, testCases2Accident, 3 /* accident */, bonusmalus.BonusMalusPublicUsage)
	assertScoreTable(t, testCases2Accident, 99 /* accident */, bonusmalus.BonusMalusPublicUsage)
}

func assertScoreTable(t *testing.T, scoreTests []scoreTestCase, accidents uint, usage bonusmalus.BonusMalusUsage) {
	for _, testCase := range scoreTests {
		t.Run(fmt.Sprintf("%d driving years %d accidents", testCase.drivingYears, accidents), func(t *testing.T) {
			bonusMalus := bonusmalus.NewBonusMalus(testCase.drivingYears, accidents, usage)

			assert.Equal(t, testCase.expected, bonusMalus.Score())
		})
	}
}
