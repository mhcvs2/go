package main

import (
	"bytes"
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func Length(head *Node) int {
	l := 0
	p := head.next
	for p != nil {
		l++
		p = p.next
	}
	return l
}

func Add(head, node *Node){
	q := head
	p := head
	for p != nil {
		q = p
		p = p.next
	}
	q.next = node
	node.next = nil
}

func String(head *Node) string {
	var res bytes.Buffer
	res.WriteString("[")

	for p := head.next; p != nil; p = p.next {
		res.WriteString(fmt.Sprintf("%d", p.data))
		if p.next != nil {
			res.WriteString(", ")
		}
	}
	res.WriteString("]")
	return res.String()
}

func AddIndex(head, node *Node, index int) bool {
	l := Length(head)
	if index > l || index < 0 {
		return false
	}
	p := head
	for i:=0; i<index; i++ {
		p = p.next
	}
	node.next = p.next
	p.next = node
	return true
}

func Delete(head *Node, index int) bool {
	l := Length(head)
	if index > l || index < 1 {
		return false
	}
	p := head
	q := head
	for i := 0; i < index; i++ {
		q = p
		p = p.next
	}
	q.next = p.next
	return true
}

func Reverse(head *Node) {
	current := head.next
	p := current
	for current.next != nil {
		p = current.next
		current.next = p.next
		p.next = head.next
		head.next = p
	}
}

func SortUp(head *Node) {
	for p := head.next; p !=nil; p = p.next {
		for q := p; q != nil; q = q.next {
			if q.data < p.data {
				t := q.data
				q.data = p.data
				p.data = t
			}
		}
	}
}

func main() {
	ll := &Node{}
	Add(ll, &Node{data:1})
	Add(ll, &Node{data:2})
	Add(ll, &Node{data:3})
	Add(ll, &Node{data:4})
	Add(ll, &Node{data:5})
	fmt.Println(Length(ll))
	fmt.Println(String(ll))
	AddIndex(ll, &Node{data:999}, 1)
	fmt.Println(String(ll))
	AddIndex(ll, &Node{data:888}, 0)
	fmt.Println(String(ll))
	AddIndex(ll, &Node{data:777}, 7)
	fmt.Println(String(ll))
	Delete(ll, 8)
	fmt.Println(String(ll))
	Delete(ll, 1)
	fmt.Println(String(ll))
	Delete(ll, 2)
	fmt.Println(String(ll))
	Reverse(ll)
	fmt.Println(String(ll))
	SortUp(ll)
	fmt.Println(String(ll))
}
