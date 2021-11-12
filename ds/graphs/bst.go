package main

import "fmt"

type node struct {
	value int
	left  *node
	right *node
}
type bst struct {
	root *node
}

func (b *bst) insert(value int) {
	if b.root == nil {
		b.root = new(node)
		b.root.value = value
	} else {
		var tmp *node = b.root
		var prev *node = b.root
		for tmp != nil {
			if value < tmp.value {
				prev = tmp
				tmp = tmp.left
			} else {
				prev = tmp
				tmp = tmp.right
			}
		}
		if prev.value > value {
			prev.left = new(node)
			prev.left.value = value
		} else {
			prev.right = new(node)
			prev.right.value = value
		}
	}
}
func (b *bst) remove(key int) {
	if b.root == nil {
		return
	}
	if b.root.value == key {
		if b.root.left != nil && b.root.right != nil {
			var curr *node
			curr = b.root.left
			for curr.right != nil {
				curr = curr.right
			}
			b.root.right.left = curr
			b.root = b.root.right
		} else {
			if b.root.left != nil {
				b.root = b.root.left
			} else if b.root.right != nil {
				b.root = b.root.right
			}
		}
	} else {
		// finding node
		var curr *node = b.root
		var parent *node = b.root
		for curr.value != key {
			parent = curr
			if key > curr.value {
				curr = curr.right
			} else {
				curr = curr.left
			}
		}
		if curr.left == nil && curr.right == nil { //  node has no child
			if parent.value > curr.value {
				parent.left = nil
			} else {
				parent.right = nil
			}
		} else if curr.left != nil && curr.right != nil { // node has both childs
			if parent.value > curr.value {
				parent.left = curr.left
				curr.left.right = curr.right
			} else {
				parent.right = curr.right
				curr.right.left = curr.left
			}
		} else { // node has atleast one child
			if curr.left != nil {
				if parent.value > curr.value {
					parent.left = curr.left
				}
			} else if curr.right != nil {
				if parent.value < curr.value {
					parent.right = curr.right
				}
			}
		}
	}
}
func (b *bst) inorder(curr *node) {
	if curr == nil {
		return
	} else {
		b.inorder(curr.left)
		fmt.Println(curr.value)
		b.inorder(curr.right)
	}
}
func (b *bst) print_inorder() {
	b.inorder(b.root)
}

func main() {
	b := bst{}
	b.insert(10)
	b.insert(9)
	b.insert(33)
	b.insert(13)
	b.insert(99)
	b.insert(332)
	b.print_inorder()
	fmt.Println()
	b.remove(99)
	b.remove(332)
	b.remove(9)
	b.remove(10)
	b.insert(99898)
	b.remove(33)
	b.print_inorder()
}
