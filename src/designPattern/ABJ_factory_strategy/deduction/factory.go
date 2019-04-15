package deduction

type StragegyMan int

const (
	STEADY_DEDUCTION StragegyMan = iota
	FREE_DEDUCTION
)

func GetDeduction(man StragegyMan) IDeduction {
	switch man {
	case STEADY_DEDUCTION:
		return new(SteadyDeduction)
	case FREE_DEDUCTION:
		return new(FreeDeduction)
	default:
		return new(SteadyDeduction)
	}
}
