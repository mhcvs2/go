package main

import (
	"bytes"
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func Create() *Node {
	head := Node{}
	return &head
}

func (head *Node) Length() int {
	l := 0
	p := head.next
	for {
		if p == nil { break }
		l++
		p = p.next
	}
	return l
}

func (head *Node) Length2() int {
	l := 0
	for p := head.next; p != nil; p = p.next {
		l++
	}
	return l
}

func (head *Node) AddNode(node *Node) bool {
	p := head.next
	q := head
	for {
		if p == nil { break }
		q = p
		p = p.next
	}
	q.next = node
	node.next = nil
	return true
}

func (head *Node) AddNode2(node *Node) bool {
	q := head
	for p := head.next; p != nil; p = p.next {
		q = p
	}
	q.next = node
	node.next = nil
	return true
}

func (head *Node) AddNodeIndex(node *Node, index int) bool {
	if head.Length() < index || index < 0 {
		return false
	}
	p := head
	for i:=0; i<index; i++ {
		p = p.next
	}
	t := p.next
	p.next = node
	node.next = t
	return true
}

func (head *Node) GetByIndex(index int) *Node {
	if head.Length() < index || index < 1 {
		return nil
	}
	p := head
	for i:=0; i<index; i++ {
		p = p.next
	}
	return p
}

func (head *Node) GetValue(index int) int {
	node := head.GetByIndex(index)
	if node == nil {
		return -1
	}
	return node.data
}

func (head *Node) updateValue(index, newValue int) bool {
	node := head.GetByIndex(index)
	if node == nil {
		return false
	}
	node.data = newValue
	return true
}

func (head *Node) DeleteNode(index int) bool {
	if head.Length() < index || index < 1 {
		return false
	}
	q := head
	p := head
	for i:=0; i<index; i++ {
		q = p
		p = p.next
	}
	t := p.next
	q.next = t
	return true
}

func (head *Node) Clear() bool {
	head.next = nil
	return true
}

func (head *Node) String() string {
	var res bytes.Buffer
	res.WriteString("[")
	p := head.next
	for {
		if p == nil { break }
		res.WriteString(fmt.Sprintf("%d", p.data))
		p = p.next
		if p != nil {
			res.WriteString(", ")
		}
	}
	res.WriteString("]")
	return res.String()
}

func (head *Node) SortUp() {
	l := head.Length()
	for i:=1; i<=l; i++ {
		for j:=i; j<=l; j++ {
			di := head.GetValue(i)
			dj := head.GetValue(j)
			if dj < di {
				head.updateValue(i, dj)
				head.updateValue(j, di)
			}
		}
	}
}

func (head *Node) SortDown() {
	l := head.Length()
	for i:=1; i<=l; i++ {
		for j:=i; j<=l; j++ {
			di := head.GetValue(i)
			dj := head.GetValue(j)
			if dj > di {
				head.updateValue(i, dj)
				head.updateValue(j, di)
			}
		}
	}
}

func (head *Node) SortUp2() {
	for p := head.next; p != nil; p = p.next {
		for q := p; q != nil; q = q.next {
			if q.data < p.data {
				t := p.data
				p.data = q.data
				q.data = t
			}
		}
	}
}

func main() {
	list := Create()
	list.AddNode2(&Node{data:3})
	list.AddNode2(&Node{data:4})
	list.AddNode2(&Node{data:5})
	list.AddNode2(&Node{data:6})
	fmt.Println(list.String())
	fmt.Println(list.DeleteNode(2))
	fmt.Println(list.String())
	fmt.Println(list.Length2())
	list.AddNodeIndex(&Node{data:999}, 1)
	fmt.Println(list.String())
	list.AddNodeIndex(&Node{data:888}, 0)
	fmt.Println(list.String())
	list.AddNodeIndex(&Node{data:777}, 5)
	fmt.Println(list.String())
	fmt.Println(list.GetValue(1))
	list.updateValue(2, 666)
	fmt.Println(list.String())
	list.SortUp2()
	fmt.Println(list.String())
	list.Clear()
	fmt.Println(list.String())
}