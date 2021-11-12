package main

import "fmt"

type Node struct {
	value int
	next  *Node
	prev  *Node
}
type stack struct {
	root *Node
	top  *Node
}

func (s *stack) push(value int) {
	if s.root == nil {
		s.root = new(Node)
		s.root.value = value
		s.top = s.root
	} else {
		s.top.next = new(Node)
		s.top.next.prev = s.top
		s.top = s.top.next
		s.top.value = value
	}
}
func (s *stack) pop() (int, bool) {
	if s.top == nil {
		return -1, false
	} else {
		if s.root == s.top {
			val := s.top.value
			s.root, s.top = nil, nil
			return val, true
		} else {
			val := s.top.value
			s.top.prev.next = nil
			s.top = s.top.prev
			return val, true
		}
	}
}

func (s *stack) print() {
	var tmp *Node
	for tmp = s.root; tmp != nil; tmp = tmp.next {
		fmt.Println(tmp.value)
	}
}

func main() {
	stk := stack{}
	stk.push(1)
	stk.push(2)
	stk.push(3)
	stk.print()

	stk.pop()
	stk.pop()
	stk.pop()
	stk.pop()

	stk.push(9999999988888)
	stk.push(9999999988888111)

	stk.print()

}
