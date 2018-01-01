package main

import (
	"fmt"
)

type Node struct {
	data int
	lnode *Node
	rnode *Node
}

func Create(root int) *Node {
	return &Node{data:root}
}

func (bt *Node) Add(child ...int) {
	l := len(child)
	if l >= 1 {
		bt.lnode = &Node{data:child[0]}
	}
	if l > 1 {
		bt.rnode = &Node{data:child[1]}
	}
}

func (node *Node) PreOrder() (res []int ){
	stack := make([]*Node, 100)
	top := 0
	t := node

	for t != nil || top > 0 {
		for t != nil {
			res = append(res, t.data)
			stack[top] = t
			top++
			t = t.lnode
		}
		if top > 0 {
			t = stack[top-1].rnode
			top--
		}
	}

	return res
}

func (node *Node) InOrder() (res []int) {
	stack := make([]*Node, 100)
	top := 0
	t := node

	for t != nil || top > 0 {
		for t != nil {
			stack[top] = t
			top++
			t = t.lnode
		}
		if top > 0 {
			t = stack[top-1]
			top--
			res = append(res, t.data)
			t = t.rnode
		}
	}
	return res
}

func (node *Node) PostOrder() (res []int) {
	stack := make([]*Node, 100)
	flagStack := make([]bool, 100)
	top, ftop := 0, 0
	t := node

	for t != nil || top > 0 {
		for t != nil {
			stack[top] = t
			top++
			flagStack[ftop] = false
			ftop++
			t = t.lnode
		}
		//for flag := flagStack[ftop-1]; ftop >0 && flag; {
		//	res = append(res, stack[top-1].data)
		//	top--
		//	ftop--
		//	if ftop > 0 {
		//		flag = flagStack[ftop-1]
		//	}
		//}
		for ;ftop >0 && flagStack[ftop-1]; ftop-- {
			res = append(res, stack[top-1].data)
			top--
		}
		if top > 0 {
			flagStack[ftop-1] = true
			t = stack[top-1].rnode
		}
	}
	return res
}

func main() {
	bt := Create(1)
	bt.Add(2, 3)
	bt.lnode.Add(4, 5)
	bt.rnode.Add(6, 7)
	fmt.Println(bt.PreOrder())
	fmt.Println(bt.InOrder())
	fmt.Println(bt.PostOrder())
}