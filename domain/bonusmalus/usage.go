package bonusmalus

type BonusMalusUsage string

var (
	BonusMalusPrivateUsage       BonusMalusUsage = "PRIVATE"
	BonusMalusProfessionnalUsage BonusMalusUsage = "PROFESIONNAL"
)

func (b BonusMalusUsage) IsValid() bool {
	switch b {
	case BonusMalusPrivateUsage, BonusMalusProfessionnalUsage:
		return true
	}

	return false
}
