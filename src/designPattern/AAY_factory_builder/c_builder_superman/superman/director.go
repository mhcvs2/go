package superman

// 如此折行...
func GetAdultSuperMan() *SuperMan {
	return NewSuperManBuilder().
		AndBody("strong").
		AndSpecialSmbol("s").
		AndSpecialTalent("fly").
		Build()
}


func GetChildSuperMan() *SuperMan {
	return NewSuperManBuilder().
		AndBody("strong2").
		AndSpecialSmbol("s2").
		AndSpecialTalent("fly2").
		Build()
}