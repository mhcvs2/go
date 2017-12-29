package main

import (
	"basicAlgorithm/structure"
	"fmt"
)

type Node struct {
	data int
	lnode *Node
	rnode *Node
}

func NewBinaryTree(data int) *Node {
	return &Node{data:data}
}

func (node *Node) PreOrder() (res []int) {
	t := node
	stack := structure.NewStack(7)
	for t!=nil || !stack.Empty() {
		for t!=nil{
			res = append(res, t.data)
			stack.Push(t)
			t = t.lnode
		}
		if !stack.Empty() {
			v, _ := stack.Pop()
			t = v.(*Node)
			t = t.rnode
		}
	}
	return res
}

func (node *Node) InOrder() (res []int) {
	t := node
	stack := structure.NewStack(7)
	for t != nil || !stack.Empty() {
		for t != nil {
			stack.Push(t)
			t = t.lnode
		}
		if !stack.Empty() {
			v, _ := stack.Pop()
			t = v.(*Node)
			res = append(res, t.data)
			t = t.rnode
		}
	}
	return res
}

func (node *Node) PostOrder() (res []int) {
	t := node
	stack := structure.NewStack(7)
	s := structure.NewStack(7)
	for t != nil || !stack.Empty() {
		for t != nil {
			stack.Push(t)
			s.Push(false)
			t = t.lnode
		}
		for flag, _ := s.Top(); !stack.Empty() && flag.(bool); {
			s.Pop()
			v, _ := stack.Pop()
			res = append(res, v.(*Node).data)
			flag, _ = s.Top()
		}
		if !stack.Empty() {
			s.Pop()
			s.Push(true)
			v, _ := stack.Top()
			t = v.(*Node)
			t = t.rnode
		}
	}
	return res
}

func main() {
	bt := NewBinaryTree(1)
	bt.lnode = &Node{data:2}
	bt.rnode = &Node{data:3}
	bt.lnode.lnode = &Node{data:4}
	bt.lnode.rnode = &Node{data:5}
	bt.rnode.lnode = &Node{data:6}
	bt.rnode.rnode = &Node{data:7}
	fmt.Println(bt.PreOrder())
	fmt.Println(bt.InOrder())
	fmt.Println(bt.PostOrder())
}
//[1 2 4 5 3 6 7]
//[4 2 5 1 6 3 7]
//[4 5 2 6 7 3 1]