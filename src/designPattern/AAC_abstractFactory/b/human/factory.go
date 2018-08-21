package human

type HumanType int

const (
	BLACKHUMAN HumanType = iota
	YELLOWHUMAN
	WHITEHUMAN
	FEMALEBLACKHUMAN
	MALEBLACKHUMAN
)

func CreateHuman(t HumanType) Human {
	switch t {
	case BLACKHUMAN:
		return &BlackHuman{}
	case YELLOWHUMAN:
		return &YellowHuman{}
	case WHITEHUMAN:
		return &WhiteHuman{}
	default:
		return nil
	}
}

func CreateSexHuman(t HumanType) SexHuman {
	switch t {
	case FEMALEBLACKHUMAN:
		return &FemaleBlackHuman{}
	case MALEBLACKHUMAN:
		return &MaleBlackHuman{}
	default:
		return nil
	}
}
