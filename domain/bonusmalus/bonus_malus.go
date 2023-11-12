package bonusmalus

type BonusMalusUsage string

var (
	BonusMalusPrivateUsage BonusMalusUsage = "PRIVATE"
	BonusMalusPublicUsage  BonusMalusUsage = "PUBLIC"
)

type bonusMalus struct {
}

func NewBonusMalus(drivingYears, accidents uint, usage BonusMalusUsage) bonusMalus {
	return bonusMalus{}
}

func (b bonusMalus) Score() int
