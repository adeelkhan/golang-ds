package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type slinkedlist struct {
	root *Node
	curr *Node
}

func (s *slinkedlist) Init() {
	s.root = nil
}
func (s *slinkedlist) insert(value int) {
	if s.root == nil {
		s.root = new(Node)
		s.curr = s.root
		s.root.value, s.root.next = value, nil
	} else {
		s.curr.next = new(Node)
		s.curr = s.curr.next
		s.curr.value, s.curr.next = value, nil
	}
}
func (s *slinkedlist) remove(key int) bool {
	if s.root == nil { // nothing
		return false
	}
	if s.root == s.curr && s.root.value == key { // single value
		s.root, s.curr = nil, nil
		return true
	}
	if s.root.value == key { // root value
		s.root = s.root.next
		return true
	}

	// finding key position
	var tmp *Node = nil
	for tmp = s.root; tmp.next.value != key; tmp = tmp.next {
	}
	if tmp == nil {
		return false
	} else if tmp.next == s.curr {
		tmp.next = tmp.next.next
		s.curr = tmp
	} else {
		tmp.next = tmp.next.next
	}
	return true
}
func (s *slinkedlist) print() {
	for tmp := s.root; tmp != nil; tmp = tmp.next {
		fmt.Println(tmp.value)
	}
}

func main() {
	slink := slinkedlist{}
	slink.insert(10)
	slink.insert(12)
	slink.insert(13)
	slink.print()

	slink.remove(10)
	slink.remove(13)
	slink.print()
}
