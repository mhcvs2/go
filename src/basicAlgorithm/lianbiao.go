package main

import (
	"bytes"
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func Create(data int) *Node {
	head := Node{data:data}
	return &head
}

func (head *Node) addNode(node *Node) bool {
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

func (head *Node) deleteNode(index int) bool {
	p := head.next
	l := 0
	for {
		if p == nil { break }
		l++
		p = p.next
	}
	if l < index {
		return false
	} else {
		q := head
		p = head
		for i:=0; i<index; i++ {
			q = p
			p = p.next
		}
		t := p.next
		q.next = t
		return true
	}
}

func (head *Node) clear() bool {
	p := head
	p.next = nil
	return true
}

func (head *Node) String() string {
	var res bytes.Buffer
	p := head
	for {
		if p == nil { break }
		res.WriteString(fmt.Sprintf("%d ", p.data))
		p = p.next
	}
	return res.String()
}

func main() {
	list := Create(9)
	list.addNode(&Node{data:3})
	list.addNode(&Node{data:4})
	list.addNode(&Node{data:5})
	list.addNode(&Node{data:6})
	fmt.Println(list.String())
	fmt.Println(list.deleteNode(2))
	fmt.Println(list.String())
}