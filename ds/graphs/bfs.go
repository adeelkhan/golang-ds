package main

import (
	"fmt"

	"github.com/golang-collections/go-datastructures/queue"
)

type graph struct {
	g    [][]int
	size uint64
}

func (gr *graph) Init(size uint64) {
	gr.size = size
	gr.g = make([][]int, gr.size)
	var i uint64
	for i = 0; i < gr.size; i++ {
		gr.g[i] = make([]int, gr.size)
	}
}

func (gr *graph) SetConnection(u, v uint64) {
	if u < gr.size && v < gr.size {
		gr.g[u][v] = 1
	}
}

func (gr *graph) IsConnected(u, v uint64) bool {
	if u < gr.size && v < gr.size {
		if gr.g[u][v] == 1 {
			return true
		} else {
			return false
		}
	}
	return false
}
func (gr *graph) print() {
	var i, j uint64
	for i = 0; i < gr.size; i++ {
		for j = 0; j < gr.size; j++ {
			fmt.Printf("%d", gr.g[i][j])
		}
		fmt.Println()
	}
}

func bfs_walk(q *queue.RingBuffer, gr *graph, visit []uint64) {
	for q.Len() != 0 {
		u, error := q.Get()
		if error != nil {
			return
		}

		src, ok := u.(uint64)
		fmt.Println(src)
		var dest uint64
		if ok {
			for dest = 0; dest < gr.size; dest++ {
				if visit[dest] != 1 && gr.IsConnected(src, dest) == true {
					q.Put(dest)
					visit[dest] = 1
				}
			}
			bfs_walk(q, gr, visit)
		}
	}
}
func bfs(gr *graph) {
	que := queue.NewRingBuffer(gr.size)
	visit := make([]uint64, gr.size)
	gr.print()

	var src uint64 = 1
	visit[src] = 1
	que.Put(src)

	bfs_walk(que, gr, visit)
}

func main() {
	g := graph{}
	g.Init(6)

	g.SetConnection(1, 2)
	g.SetConnection(1, 3)
	g.SetConnection(2, 3)
	g.SetConnection(2, 4)
	g.SetConnection(3, 5)
	g.SetConnection(4, 5)
	bfs(&g)

}
