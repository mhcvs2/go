package human

type IFactory interface {
	CreateBlackHuman() SexHuman
	CreateYellowHuman() SexHuman
	CreateWhiteHuman() SexHuman
}

type MaleFactory struct {
}

func (f *MaleFactory) CreateBlackHuman() SexHuman {
	return new(MaleBlackHuman)
}

func (f *MaleFactory) CreateYellowHuman() SexHuman {
	return new(MaleBlackHuman)   //偷懒了
}

func (f *MaleFactory) CreateWhileHuman() SexHuman {
	return new(MaleBlackHuman)
}

type FemaleFactory struct {
}

func (f *FemaleFactory) CreateBlackHuman() SexHuman {
	return new(FemaleBlackHuman)
}

func (f *FemaleFactory) CreateYellowHuman() SexHuman {
	return new(FemaleBlackHuman)   //偷懒了
}

func (f *FemaleFactory) CreateWhileHuman() SexHuman {
	return new(FemaleBlackHuman)
}
