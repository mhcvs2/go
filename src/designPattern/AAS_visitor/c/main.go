package main

import "designPattern/AAS_visitor/c/visitor"

func main() {
	actor := visitor.OldActor{}
	role := visitor.KungFuRole{}
	role.Accept(actor)
}
//不喜欢功夫角色
