package visitor

import "fmt"

type IActor interface {
	Act(role Role)
}

type Actor struct {
}

func (a Actor) Act(role Role) {
	fmt.Println("演员可以扮演任何角色")
}

type YoungActor struct {
	Actor
}

func (a YoungActor) Act(role Role) {
	fmt.Println("最喜欢功夫角色")
}

type OldActor struct {
	Actor
}

func (a OldActor) Act(role Role) {
	fmt.Println("不喜欢功夫角色")
}