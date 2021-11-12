package main

import "fmt"

type Node struct {
	value int
	next  *Node
	prev  *Node
}
type queueL struct {
	front *Node
	rear  *Node
}

func (q *queueL) enqueue(value int) {
	if q.rear == nil {
		q.rear = new(Node)
		q.front = q.rear
		q.rear.value = value
	} else {
		q.rear.next = new(Node)
		q.rear.next.prev = q.rear
		q.rear = q.rear.next
		q.rear.value = value
	}
}
func (q *queueL) dequeue() (int, bool) {
	if q.front == nil {
		return -1, false
	} else if q.front == q.rear {
		val := q.front.value
		q.front, q.rear = nil, nil
		return val, true
	} else {
		val := q.front.value
		q.front.next.prev = nil
		q.front = q.front.next
		return val, true
	}
}
func (q *queueL) print() {
	if q.front == nil || q.rear == nil {
		return
	}
	var tmp *Node
	for tmp = q.front; tmp != q.rear.next; tmp = tmp.next {
		fmt.Println(tmp.value)
	}
}

func main() {
	q := queueL{}
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.print()

	fmt.Println()
	q.dequeue()
	q.dequeue()
	q.dequeue()
	q.dequeue()
	q.enqueue(3)
	q.print()

}
