package main

import "fmt"

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type dlinkedlist struct {
	root *Node
}

func (d *dlinkedlist) Init() {
	d.root = nil
}
func (d *dlinkedlist) Insert(value int) {
	if d.root == nil {
		d.root = new(Node)
		d.root.value = value
		d.root.next, d.root.prev = nil, nil
	} else {
		var tmp *Node
		tmp = new(Node)
		tmp.value = value

		d.root.prev = tmp
		tmp.next, tmp.prev = d.root, nil
		d.root = tmp
	}
}
func (d *dlinkedlist) Remove(key int) {
	var root *Node = d.root
	if root == nil {
		return
	}
	var tmp *Node
	// finding key
	for tmp = root; tmp.value != key; tmp = tmp.next {
	}

	if tmp == root &&
		tmp.next == nil &&
		tmp.prev == nil {
		d.root = nil
		return
	}

	if tmp.prev != nil && tmp.next != nil {
		tmp.prev.next = tmp.next
		tmp.next.prev = tmp.prev
	} else {
		if tmp.next != nil {
			tmp.next.prev = tmp.prev
		} else if tmp.prev != nil {
			tmp.prev.next = tmp.next
		}
	}
	if tmp == root && root.prev != nil { // setting root
		d.root = root.prev
	}
	if tmp == root && root.next != nil {
		d.root = root.next
	}
}

func (d *dlinkedlist) print() {
	for tmp := d.root; tmp != nil; tmp = tmp.next {
		fmt.Println(tmp.value)
	}
}

func main() {
	dlink := dlinkedlist{}
	dlink.Insert(10)
	dlink.Insert(12)
	dlink.Insert(13)
	dlink.print()

	dlink.Remove(10)
	dlink.Remove(13)
	dlink.Remove(12)
	dlink.print()
}
