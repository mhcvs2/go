package main

import "designPattern/adapter/user"

func main() {
	var ui user.IUserInfo
	ui = new(user.UserInfo)
	ui.GetName()
	ui.GetNumber()

	ui = user.NewOuterUserInfo()
	ui.GetName()
	ui.GetNumber()
}

//get user name
//get user number
//get outer user name
//get outer user number