package human

type HumanType int

const (
	BLACKHUMAN HumanType = iota
	YELLOWHUMAN
	WHITEHUMAN
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

