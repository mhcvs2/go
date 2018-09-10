package user

import "fmt"

type OuterUserInfo struct {
	OuterUser
	baseInfo map[string]string
	homeInfo map[string]string
}

func NewOuterUserInfo() *OuterUserInfo {
	res := &OuterUserInfo{}
	res.baseInfo = res.OuterUser.GetBaseInfo()
	res.homeInfo = res.OuterUser.GetHomeInfo()
	return res
}

func (u *OuterUserInfo) GetName() string {
	fmt.Println("get outer user name")
	return u.baseInfo["name"]
}

func (u *OuterUserInfo) GetNumber() string {
	fmt.Println("get outer user number")
	return u.homeInfo["number"]
}