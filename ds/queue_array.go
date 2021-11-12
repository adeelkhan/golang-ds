package main

import "fmt"

type queue struct {
	data  []int64
	qsize int
	front int
	rear  int
}

func (q *queue) Init(size int) {
	q.qsize = size
	q.data = make([]int64, q.qsize)
	q.front, q.rear = -1, -1
}

func (q *queue) Enqueue(value int64) bool {
	if q.IsFull() != true {
		q.data[q.rear+1] = value
		q.rear += 1

		return true
	} else {
		return false
	}
}
func (q *queue) Dequeue() (int64, bool) {
	if q.IsEmpty() != true {
		value := q.data[q.front+1]
		// left shift values
		for i := 1; i <= q.rear; i++ {
			q.data[i-1] = q.data[i]
			q.data[i] = 0
		}
		q.rear -= 1
		return value, true
	} else {
		return -1, false
	}

}
func (q *queue) IsFull() bool {
	if q.rear < q.qsize {
		return false
	}
	return true
}
func (q *queue) IsEmpty() bool {
	if q.front == q.rear {
		return true
	}
	return false
}

func main() {
	q := queue{}
	q.Init(10)
	q.Enqueue(10)
	q.Enqueue(22)
	fmt.Println(q.data)
	q.Dequeue()
	fmt.Println(q.data)
	q.Enqueue(9899)
	q.Enqueue(9899)
	fmt.Println(q.data)
	q.Dequeue()
	fmt.Println(q.data)

}
