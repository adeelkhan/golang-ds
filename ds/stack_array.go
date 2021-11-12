package main

import "fmt"

type stack struct {
	data  []int64
	ssize int
	top   int
}

func (v *stack) Init(ssize int) {
	v.ssize = ssize
	v.data = make([]int64, v.ssize)
	v.top = -1
}

func (v *stack) Push(value int64) {
	if v.IsFull() != true {
		v.top += 1
		v.data[v.top] = value
	} else {
		fmt.Println("Stack is full")
	}
}
func (v *stack) Pop() (int64, bool) {
	if v.IsEmpty() != true {
		value := v.data[v.top]
		v.top -= 1
		return value, true
	} else {
		fmt.Println("Stack is empty")
		return -1, false
	}
}
func (v *stack) IsEmpty() bool {
	if v.top == -1 {
		return true
	} else {
		return false
	}
}
func (v *stack) IsFull() bool {
	if v.top < v.ssize-1 {
		return false
	} else {
		return true
	}
}

func main() {
	stk := stack{}
	stk.Init(10)

	stk.Push(12)
	stk.Push(13)
	fmt.Println(stk.data)

	stk.Pop()
	stk.Push(99)
	fmt.Println(stk.data)

}
