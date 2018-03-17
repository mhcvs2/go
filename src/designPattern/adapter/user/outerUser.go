package user

type IOuterUser interface {
	GetBaseInfo() map[string]string
	GetHomeInfo() map[string]string
}

type OuterUser struct {}

func (u *OuterUser) GetBaseInfo() map[string]string {
	res := make(map[string]string)
	res["name"] = "aaa"
	res["sex"] = "nama"
	return res
}

func (u *OuterUser) GetHomeInfo() map[string]string {
	res := make(map[string]string)
	res["number"] = "123"
	res["address"] = "china"
	return res
}
