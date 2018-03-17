package user

import "fmt"

type IUserInfo interface {
	GetName() string
	GetNumber() string
}

type UserInfo struct {}

func (u *UserInfo) GetName() string {
	fmt.Println("get user name")
	return ""
}

func (u *UserInfo) GetNumber() string {
	fmt.Println("get user number")
	return ""
}