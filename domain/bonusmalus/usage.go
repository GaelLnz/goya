package bonusmalus

type BonusMalusUsage string

var (
	BonusMalusPrivateUsage       BonusMalusUsage = "PRIVATE"
	BonusMalusProfessionnalUsage BonusMalusUsage = "PRO"
)

func (b BonusMalusUsage) IsValid() bool {
	switch b {
	case BonusMalusPrivateUsage, BonusMalusProfessionnalUsage:
		return true

	}

	return false
}
