package bonusmalus

type BonusMalusUsage string

var (
	BonusMalusPrivateUsage BonusMalusUsage = "PRIVATE"
	BonusMalusPublicUsage  BonusMalusUsage = "PUBLIC"
)

func (b BonusMalusUsage) IsValid() bool {
	switch b {
	case BonusMalusPrivateUsage, BonusMalusPublicUsage:
		return true

	}

	return false
}
