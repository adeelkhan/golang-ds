package main

import "fmt"

type pq struct {
	data []int
	size int
	elem int
}

func (p *pq) swap(a *int, b *int) {
	*a, *b = *b, *a
}

func (p *pq) siftup(i int) {
	// going up
	for p.data[(i-1)/2] < p.data[i] {
		p.swap(&p.data[(i-1)/2], &p.data[i])
		i = (i - 1) / 2
	}
}
func (p *pq) siftdown(i int) {
	left := 2*i + 1
	right := 2*i + 2
	if p.data[left] > p.data[right] {
		if p.data[i] < p.data[left] {
			p.swap(&p.data[i], &p.data[left])
		}
	} else {
		if p.data[i] < p.data[right] {
			p.swap(&p.data[i], &p.data[right])
		}
	}
}

func (p *pq) Init(size int) {
	p.size = size
	p.data = make([]int, p.size)
	p.elem = 0
}
func (p *pq) Insert(value int) {
	if p.elem <= p.size {
		p.data[p.elem] = value
		p.elem++
		p.siftup(p.elem - 1)
	}
}
func (p *pq) Remove() int {
	val := p.data[0] // removing root
	p.swap(&p.data[0], &p.data[p.elem-1])
	p.elem--

	// fixing tree
	for i := 0; i < (p.elem-1)/2; i++ {
		p.siftdown(i)
	}
	fmt.Println(val)
	return val
}
func (p *pq) Print() {
	fmt.Println(p.data)
	fmt.Println("elem:", p.elem)
}

func main() {
	p := pq{}
	p.Init(10)
	p.Insert(1)
	p.Insert(2)
	p.Insert(3)
	p.Insert(4)
	p.Insert(55)
	p.Insert(155)
	
	p.Remove()
	p.Print()
	p.Remove()
	p.Print()
	p.Remove()
	p.Print()
	p.Remove()
	p.Print()
	p.Insert(99)
	p.Print()
}
